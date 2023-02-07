package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"entgo.io/ent/dialect/sql"
	"github.com/google/go-github/v50/github"
	"github.com/kataras/go-sessions/v3"
	"github.com/labstack/echo/v4"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"github.com/trim21/errgo"

	"github-bot/ent"
	"github-bot/ent/pulls"
	"github-bot/ent/user"
)

type PRHandle struct {
	logger zerolog.Logger
	ent    *ent.Client

	// client for github app
	app githubapp.ClientCreator
}

const githubCheckRunName = "bangumi contributors"

func (h PRHandle) Index(c echo.Context) error {
	s := session.Start(c.Response(), c.Request())

	if c.QueryParams().Has("debug") {
		return c.JSON(http.StatusOK, s.GetAll())
	}

	githubId := s.GetIntDefault("github_id", 0)

	var html string
	if githubId == 0 {
		return c.HTML(http.StatusOK, `<p> github 未链接，请认证 <a href="/oauth/github">github oauth</a> </p>`)
	}

	html += fmt.Sprintf(`<p> github id %d </p>`, githubId)

	bangumiId := s.GetIntDefault("bangumi_id", 0)
	if bangumiId == 0 {
		return c.HTML(http.StatusOK, `<p> bangumi 未链接，请认证 <a href="/oauth/bangumi">bangumi oauth</a> </p>`)
	}

	html += fmt.Sprintf(`<p> bangumi id %d </p>`, bangumiId)

	html += "<h1>已完成</h1>"

	if err := h.afterOauth(c.Request().Context(), s); err != nil {
		return errgo.Trace(err)
	}

	return c.HTML(http.StatusOK, html)
}

var repoToIgnore = []string{"dev-docs", "dev-env", "issue", "api", "scripts"}

var webhookSecret = []byte(os.Getenv("GITHUB_APP_WEBHOOK_SECRET"))

func verifySign(body []byte, header string) bool {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, webhookSecret)

	// Write Data to it
	h.Write(body)

	// Get result and encode as hexadecimal string
	sha := "sha256=" + hex.EncodeToString(h.Sum(nil))

	return hmac.Equal([]byte(sha), []byte(header))
}

func (h PRHandle) handlePullRequest(c echo.Context) error {
	ctx := c.Request().Context()
	var payload github.PullRequestEvent

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return errgo.Trace(err)
	}

	if !verifySign(body, c.Request().Header.Get(github.SHA256SignatureHeader)) {
		return c.String(http.StatusBadRequest, "Signatures didn't match!")
	}

	if err := json.Unmarshal(body, &payload); err != nil {
		return errgo.Trace(err)
	}

	pr := payload.PullRequest

	h.logger.Info().
		Str("action", payload.GetAction()).
		Str("repo", payload.GetRepo().GetFullName()).
		Msg("new pull webhook")

	if pr.User.GetType() == "Bot" || pr.User.GetID() == 88366224 {
		// https://api.github.com/users/Trim21-bot
		h.logger.Info().Msg("ignore bot pr")
		return nil
	}

	repo := pr.Base.Repo.GetName()
	if lo.Contains(repoToIgnore, repo) || repo == "" {
		h.logger.Info().Str("repo", repo).Msg("skip non-code repo")
	}

	if err := h.handle(ctx, payload); err != nil {
		return errgo.Trace(err)
	}

	if err := h.checkSuite(ctx, payload); err != nil {
		return errgo.Trace(err)
	}

	return nil
}

func (h PRHandle) Handle(c echo.Context) error {
	if c.Request().Header.Get(github.EventTypeHeader) == "pull_request" {
		return h.handlePullRequest(c)
	}

	return nil
}

func (h PRHandle) checkSuite(ctx context.Context, p github.PullRequestEvent) error {
	if p.GetAction() == "close" {
		return nil
	}

	pr := p.GetPullRequest()
	repo := pr.Base.Repo

	u, err := h.ent.User.Query().Where(user.GithubID(pr.User.GetID())).Only(ctx)
	if err != nil {
		return errgo.Trace(err)
	}

	pull, err := h.ent.Pulls.Query().Where(
		pulls.Repo(repo.GetName()),
		pulls.Owner(repo.GetOwner().GetLogin()),
		pulls.NumberEQ(p.PullRequest.GetNumber()),
	).Only(ctx)
	if err != nil {
		return errgo.Trace(err)
	}

	g, err := h.app.NewInstallationClient(githubapp.GetInstallationIDFromEvent(&p))
	if err != nil {
		return errgo.Trace(err)
	}

	var output *github.CheckRunOutput
	var result string
	result = checkRunSuccess
	if u.BangumiID == 0 {
		result = checkRunActionRequired
		output = &github.CheckRunOutput{
			Title:   lo.ToPtr("请关联你的 Bangumi 账号"),
			Summary: &checkRunDetailsMessage,
		}
	}

	if pull.HeadSha == p.PullRequest.GetHead().GetSHA() {
		if pull.CheckRunID != 0 {
			_, _, err := g.Checks.UpdateCheckRun(ctx, repo.Owner.GetLogin(), repo.GetName(), pull.CheckRunID, github.UpdateCheckRunOptions{
				Name:       githubCheckRunName,
				Conclusion: &result,
				Output:     output,
			})
			if err != nil {
				return errgo.Trace(err)
			}

			err = h.ent.Pulls.UpdateOne(pull).SetCheckRunResult(result).Exec(ctx)
			if err != nil {
				return errgo.Trace(err)
			}

		}
		return nil
	}

	cr, _, err := g.Checks.CreateCheckRun(ctx, repo.Owner.GetLogin(), repo.GetName(), github.CreateCheckRunOptions{
		Name:       githubCheckRunName,
		HeadSHA:    pr.Head.GetSHA(),
		Conclusion: &result,
		Output:     output,
	})

	if err != nil {
		return errgo.Trace(err)
	}

	err = h.ent.Pulls.UpdateOne(pull).
		SetCheckRunID(cr.GetID()).
		SetHeadSha(cr.GetHeadSHA()).
		SetCheckRunResult(result).
		Exec(ctx)

	if err != nil {
		return errgo.Trace(err)
	}

	return nil
}

func (h PRHandle) handle(ctx context.Context, event github.PullRequestEvent) error {
	payload := event.GetPullRequest()
	u, err := h.ent.User.Query().Where(user.GithubID(*payload.User.ID)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return errgo.Trace(err)
		}

		u, err = h.ent.User.Create().SetGithubID(payload.User.GetID()).Save(ctx)
		if err != nil {
			return errgo.Trace(err)
		}
	}

	p, err := h.ent.Pulls.Query().Where(pulls.NumberEQ(payload.GetNumber())).Only(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			return errgo.Trace(err)
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
			return errgo.Trace(err)
		}
	}

	g, err := h.app.NewInstallationClient(githubapp.GetInstallationIDFromEvent(&event))
	if err != nil {
		return errgo.Trace(err)
	}

	if u.BangumiID == 0 && p.Comment == nil {
		c, res, err := g.Issues.CreateComment(ctx, p.Owner, p.Repo, payload.GetNumber(), &github.IssueComment{
			Body: lo.ToPtr(checkRunDetailsMessage),
		})

		if err != nil {
			b, _ := io.ReadAll(res.Body)
			logger.Err(err).Bytes("body", b).Msg("failed to create issue")
			return errgo.Trace(err)
		}

		err = h.ent.Pulls.UpdateOne(p).SetComment(*c.ID).Exec(ctx)
		if err != nil {
			return errgo.Trace(err)
		}
	}

	if payload.MergedAt != nil && p.MergedAt.IsZero() {
		if _, err := h.ent.Pulls.UpdateOne(p).SetMergedAt(payload.MergedAt.Time).Save(ctx); err != nil {
			return errgo.Trace(err)
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
		return errgo.Trace(err)
	}

	prs, err := h.ent.Pulls.Query().Where(
		pulls.HasCreatorWith(user.GithubID(githubId)),
		pulls.CheckRunResult(checkRunActionRequired),
		pulls.MergedAtIsNil(),
	).All(ctx)
	if err != nil {
		logger.Err(err).Msg("failed to get pulls")
		return errgo.Trace(err)
	}

	c, err := h.app.NewInstallationClient(installationID)
	if err != nil {
		return errgo.Trace(err)
	}

	for _, pr := range prs {
		if pr.CheckRunID != 0 {
			_, _, err := c.Checks.UpdateCheckRun(ctx, pr.Owner, pr.Repo, pr.CheckRunID, github.UpdateCheckRunOptions{
				Name:       githubCheckRunName,
				Conclusion: lo.ToPtr(checkRunSuccess),
				Output: &github.CheckRunOutput{
					Title:   lo.ToPtr(""),
					Summary: lo.ToPtr(""),
				},
			})

			if err != nil {
				return errgo.Trace(err)
			}

			if err := h.ent.Pulls.UpdateOne(pr).SetCheckRunResult(checkRunSuccess).Exec(ctx); err != nil {
				return errgo.Trace(err)
			}
		}
	}

	prs, err = h.ent.Pulls.Query().Where(
		pulls.HasCreatorWith(user.GithubID(githubId)),
		pulls.CommentNEQ(0),
		pulls.MergedAtIsNil(),
	).All(ctx)
	if err != nil {
		logger.Err(err).Msg("failed to get pulls")
		return errgo.Trace(err)
	}

	for _, pr := range prs {
		_, _, err := c.Issues.EditComment(ctx, pr.Owner, pr.Repo, *pr.Comment, &github.IssueComment{
			Body: lo.ToPtr("成功关联 bangumi ID，感谢你为 bangumi 做出的贡献"),
		})
		if err != nil {
			return errgo.Trace(err)
		}
	}

	return nil
}
