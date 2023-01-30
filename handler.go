package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/go-github/v50/github"
	"github.com/kataras/go-sessions/v3"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/samber/lo"

	"github-bot/ent"
	"github-bot/ent/pulls"
	"github-bot/ent/user"
)

type PRHandle struct {
	logger zerolog.Logger
	ent    *ent.Client
	github *github.Client
}

func (h PRHandle) Index(c echo.Context) error {
	s := session.Start(c.Response(), c.Request())

	if c.QueryParams().Has("debug") {
		return c.JSON(http.StatusOK, s.GetAll())
	}

	githubId := int(s.GetFloat64Default("github_id", 0))

	var html string
	if githubId == 0 {
		return c.HTML(http.StatusOK, `<p> github 未链接，请认证 <a href="/oauth/github">github oauth</a> </p>`)
	}

	html += fmt.Sprintf(`<p> github id %d </p>`, githubId)

	bangumiId := int(s.GetFloat64Default("bangumi_id", 0))
	if bangumiId == 0 {
		return c.HTML(http.StatusOK, `<p> bangumi 未链接，请认证 <a href="/oauth/bangumi">bangumi oauth</a> </p>`)
	}

	html += fmt.Sprintf(`<p> bangumi id %d </p>`, bangumiId)

	html += "<h1>已完成</h1>"

	if err := h.afterOauth(c.Request().Context(), s); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, html)
}

var repoToIgnore = []string{"dev-docs", "dev-env", "issue", "api", "scripts"}

var webhookSecret = []byte(os.Getenv("GITHUB_WEBHOOK_SECRET"))

func verifySign(body []byte, sign string) bool {
	if !strings.HasPrefix(sign, "sha256=") {
		return false
	}

	sign = strings.TrimPrefix(sign, "sha256=")

	rawSign, err := hex.DecodeString(sign)
	if err != nil {
		return false
	}

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, webhookSecret)

	// Write Data to it
	h.Write(body)

	// Get result and encode as hexadecimal string
	sha := h.Sum(nil)

	return subtle.ConstantTimeCompare(sha, rawSign) == 0
}

func (h PRHandle) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	var payload struct {
		Action      string             `json:"action"`
		PullRequest github.PullRequest `json:"pull_request"`
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	if !verifySign(body, c.Request().Header.Get("X-Hub-Signature-256")) {
		return c.String(http.StatusBadRequest, "Signatures didn't match!")
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}

	h.logger.Info().
		Str("action", payload.Action).
		Str("repo", payload.PullRequest.Base.Repo.GetName()).
		Msg("new pull webhook")

	if payload.PullRequest.User.GetType() == "Bot" || payload.PullRequest.User.GetID() == 88366224 {
		// https://api.github.com/users/Trim21-bot
		h.logger.Info().Msg("ignore bot pr")
		return nil
	}

	repo := payload.PullRequest.Base.Repo.GetName()
	if lo.Contains(repoToIgnore, repo) || repo == "" {
		h.logger.Info().Str("repo", repo).Msg("skip non-code repo")
	}

	return h.handle(ctx, payload.PullRequest)
}

func (h PRHandle) handle(ctx context.Context, payload github.PullRequest) error {
	u, err := h.ent.User.Query().Where(user.GithubID(*payload.User.ID)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}

		u, err = h.ent.User.Create().SetGithubID(payload.User.GetID()).Save(ctx)
		if err != nil {
			return err
		}
	}

	p, err := h.ent.Pulls.Query().Where(pulls.NumberEQ(payload.GetNumber())).Only(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}

		q := h.ent.Pulls.Create().
			SetCreatedAt(payload.CreatedAt.Time).
			SetOwner(payload.Base.Repo.Owner.GetLogin()).
			SetNumber(payload.GetNumber()).
			SetRepo(payload.Base.Repo.GetName()).
			SetCreator(u)

		if payload.MergedAt != nil {
			q = q.SetMergedAt(payload.MergedAt.Time)
		}

		p, err = q.Save(ctx)
		if err != nil {
			return err
		}
	}

	if u.BangumiID == 0 && p.Comment == nil {
		c, res, err := h.github.Issues.CreateComment(ctx, p.Owner, p.Repo, payload.GetNumber(), &github.IssueComment{
			Body: lo.ToPtr("请关联您的 bangumi ID 以方便进行贡献者统计，未关联的贡献者将不会被统计在年鉴中\n\nhttps://contributors.bgm38.com/"),
		})

		if err != nil {
			b, _ := io.ReadAll(res.Body)
			logger.Err(err).Bytes("body", b).Msg("failed to create issue")
			return err
		}

		err = h.ent.Pulls.UpdateOne(p).SetComment(*c.ID).Exec(ctx)
		if err != nil {
			return err
		}
	}

	if payload.MergedAt != nil && p.MergedAt.IsZero() {
		if _, err := h.ent.Pulls.UpdateOne(p).SetMergedAt(payload.MergedAt.Time).Save(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (h PRHandle) afterOauth(ctx context.Context, s *sessions.Session) error {
	githubId := int64(s.GetFloat64Default("github_id", 0))
	bangumiId := int64(s.GetFloat64Default("bangumi_id", 0))

	if githubId == 0 || bangumiId == 0 {
		return nil
	}

	err := h.ent.User.Create().SetGithubID(githubId).SetBangumiID(bangumiId).
		OnConflict(sql.ConflictColumns(user.FieldGithubID)).UpdateBangumiID().Exec(ctx)
	if err != nil {
		logger.Err(err).Msg("failed to save authorized user to db")
		return err
	}

	return nil
}
