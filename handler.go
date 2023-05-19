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
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-github/v52/github"
	"github.com/labstack/echo/v4"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/samber/lo"
	"github.com/trim21/errgo"

	"github-bot/config"
	"github-bot/ent"
	"github-bot/ent/pulls"
	"github-bot/ent/user"
)

type PRHandle struct {
	ent *ent.Client

	// client for GitHub app
	app githubapp.ClientCreator

	// client as App Installation
	g *github.Client
}

const githubCheckRunName = "bangumi contributors"

func (h PRHandle) Index(c echo.Context) error {
	s := session.Start(c.Response(), c.Request())

	if c.QueryParams().Has("debug") {
		return c.String(http.StatusOK, spew.Sdump(s.GetAll()))
	}

	githubId := s.GetInt64Default("github_id", 0)
	if githubId == 0 {
		return c.HTML(http.StatusOK, `<p> github 未链接，请认证 <a href="/oauth/github">github oauth</a> </p>`)
	}

	bangumiId := s.GetInt64Default("bangumi_id", 0)
	if bangumiId == 0 {
		return c.HTML(http.StatusOK, `<p> bangumi 未链接，请认证 <a href="/oauth/bangumi">bangumi oauth</a> </p>`)
	}

	html := fmt.Sprintf("<p> github id %d </p>\n<p> bangumi id %d </p>\n<h1>已完成</h1>", githubId, bangumiId)

	if err := h.afterOauth(c.Request().Context(), s); err != nil {
		return errgo.Trace(err)
	}

	return c.HTML(http.StatusOK, html)
}

var repoToIgnore = []string{"dev-docs", "dev-env", "issue", "api", "scripts"}

func verifySign(body []byte, header string) bool {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, config.WebhookSecret)

	// Write Data to it
	h.Write(body)

	// Get result and encode as hexadecimal string
	sha := "sha256=" + hex.EncodeToString(h.Sum(nil))

	return hmac.Equal([]byte(sha), []byte(header))
}

func (h PRHandle) Handle(c echo.Context) error {
	if c.Request().Header.Get(github.EventTypeHeader) != "pull_request" {
		return nil
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return errgo.Trace(err)
	}

	if !verifySign(body, c.Request().Header.Get(github.SHA256SignatureHeader)) {
		return c.String(http.StatusBadRequest, "Signatures didn't match!")
	}

	var payload github.PullRequestEvent
	if err := json.Unmarshal(body, &payload); err != nil {
		return errgo.Trace(err)
	}

	return h.handlePullRequest(c, payload)
}

func (h PRHandle) handlePullRequest(c echo.Context, payload github.PullRequestEvent) error {
	pr := payload.PullRequest

	logger.Info().
		Str("action", payload.GetAction()).
		Str("repo", payload.GetRepo().GetFullName()).
		Msg("new pull webhook")

	if pr.User.GetType() == "Bot" || strings.HasSuffix(strings.ToLower(pr.User.GetLogin()), "-bot") {
		logger.Info().Msg("ignore bot pr")
		return nil
	}

	if repo := pr.Base.Repo.GetName(); lo.Contains(repoToIgnore, repo) || repo == "" {
		logger.Info().Str("repo", repo).Msg("skip non-code repo")
		return nil
	}

	if owner := pr.Base.Repo.GetOwner().GetLogin(); owner != "bangumi" {
		logger.Info().Str("repo_owner", owner).Msg("skip non-bangumi repo")
		return nil
	}

	ctx := c.Request().Context()

	u, p, err := h.objectFromEvent(ctx, payload)
	if err != nil {
		return errgo.Trace(err)
	}

	if err := h.handle(ctx, u, p); err != nil {
		return errgo.Trace(err)
	}

	if err := h.checkSuite(ctx, u, p, payload.PullRequest); err != nil {
		return errgo.Trace(err)
	}

	return nil
}

func (h PRHandle) handle(ctx context.Context, u *ent.User, p *ent.Pulls) error {
	if u.BangumiID == 0 && p.Comment == 0 {
		c, res, err := h.g.Issues.CreateComment(ctx, p.Owner, p.Repo, p.Number, &github.IssueComment{
			Body: lo.ToPtr(checkRunDetailsMessage),
		})

		if err != nil {
			b, _ := io.ReadAll(res.Body)
			logger.Err(err).Bytes("body", b).Msg("failed to create issue")
			return errgo.Trace(err)
		}

		err = h.ent.Pulls.UpdateOne(p).SetComment(c.GetID()).Exec(ctx)
		return errgo.Trace(err)
	}

	return nil
}

func (h PRHandle) checkSuite(
	ctx context.Context,
	u *ent.User, p *ent.Pulls,
	pr *github.PullRequest,
) error {
	if pr.GetState() == "closed" {
		return nil
	}

	if p.HeadSha == pr.GetHead().GetSHA() {
		return nil
	}

	repo := pr.Base.Repo

	if u.BangumiID != 0 {
		_, _, err := h.g.Checks.CreateCheckRun(ctx, repo.Owner.GetLogin(), repo.GetName(), github.CreateCheckRunOptions{
			Name:       githubCheckRunName,
			HeadSHA:    pr.Head.GetSHA(),
			Conclusion: &checkRunSuccess,
			Output: &github.CheckRunOutput{
				Title:   lo.ToPtr("Bangumi 账号已关联"),
				Summary: &successMessage,
			},
		})
		return errgo.Trace(err)
	}

	cr, _, err := h.g.Checks.CreateCheckRun(ctx, repo.Owner.GetLogin(), repo.GetName(), github.CreateCheckRunOptions{
		Name:       githubCheckRunName,
		HeadSHA:    pr.Head.GetSHA(),
		Conclusion: &checkRunActionRequired,
		Output: &github.CheckRunOutput{
			Title:   lo.ToPtr("请关联你的 Bangumi 账号"),
			Summary: &checkRunDetailsMessage,
		},
	})

	if err != nil {
		return errgo.Trace(err)
	}

	err = h.ent.Pulls.UpdateOne(p).
		SetCheckRunID(cr.GetID()).
		SetHeadSha(pr.GetHead().GetSHA()).
		SetCheckRunResult(checkRunActionRequired).
		Exec(ctx)

	return errgo.Trace(err)
}

func (h PRHandle) objectFromEvent(ctx context.Context, event github.PullRequestEvent) (*ent.User, *ent.Pulls, error) {
	pr := event.GetPullRequest()

	u, err := ent.WithTxR(ctx, h.ent, func(tx *ent.Tx) (*ent.User, error) {
		u, err := h.ent.User.Query().Where(user.GithubID(pr.User.GetID())).Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, errgo.Trace(err)
			}

			u, err = h.ent.User.Create().SetGithubID(pr.User.GetID()).Save(ctx)
			if err != nil {
				return u, errgo.Trace(err)
			}
		}

		return u, nil
	})

	if err != nil {
		return nil, nil, errgo.Trace(err)
	}

	repo := pr.GetBase().GetRepo()

	p, err := h.ent.Pulls.Query().Where(
		pulls.Or(
			pulls.PrIDEQ(pr.GetID()),
			pulls.And(
				pulls.NumberEQ(pr.GetNumber()),
				pulls.OwnerEQ(repo.GetOwner().GetLogin()),
				pulls.RepoEQ(repo.GetName()),
			),
		),
	).Only(ctx)
	if err == nil {
		return u, p, nil
	}

	if !ent.IsNotFound(err) {
		return nil, nil, errgo.Trace(err)
	}

	q := h.ent.Pulls.Create().
		SetCreatedAt(pr.CreatedAt.Time).
		SetOwner(repo.Owner.GetLogin()).
		SetRepo(repo.GetName()).
		SetNumber(pr.GetNumber()).
		SetRepoID(repo.GetID()).
		SetPrID(pr.GetID()).
		SetCreator(u)

	if pr.MergedAt != nil {
		q = q.SetMergedAt(pr.MergedAt.Time)
	}

	p, err = q.Save(ctx)
	if err != nil {
		return u, p, errgo.Trace(err)
	}

	var mutation []func(u *ent.PullsUpdateOne) *ent.PullsUpdateOne

	if pr.MergedAt != nil && p.MergedAt.IsZero() {
		mutation = append(mutation, func(u *ent.PullsUpdateOne) *ent.PullsUpdateOne {
			return u.SetMergedAt(pr.MergedAt.Time)
		})
	}

	if p.RepoID == 0 {
		mutation = append(mutation, func(u *ent.PullsUpdateOne) *ent.PullsUpdateOne {
			return u.SetRepoID(pr.Base.Repo.GetID())
		})
	}

	if p.PrID == 0 {
		mutation = append(mutation, func(u *ent.PullsUpdateOne) *ent.PullsUpdateOne {
			return u.SetPrID(pr.GetID())
		})
	}

	if len(mutation) != 0 {
		updateOne := h.ent.Pulls.UpdateOne(p)
		for _, f := range mutation {
			updateOne = f(updateOne)
		}

		p, err = updateOne.Save(ctx)
		if err != nil {
			logger.Err(err).Msg("failed to update pulls")
			return nil, nil, errgo.Trace(err)
		}
	}

	return u, p, nil
}
