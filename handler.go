package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/v49/github"
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

func (h PRHandle) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	var payload struct {
		PullRequest github.PullRequest `json:"pull_request"`
	}
	if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return err
	}

	h.logger.Debug().Interface("payload", payload).Msg("new event")

	return h.handle(ctx, payload.PullRequest)
}

func (h PRHandle) handle(ctx context.Context, payload github.PullRequest) error {
	u, err := h.ent.User.Query().Where(user.GithubID(*payload.User.ID)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}

		u, err = h.ent.User.Create().SetGithubID(*payload.User.ID).Save(ctx)
		if err != nil {
			return err
		}
	}

	p, err := h.ent.Pulls.Query().Where(pulls.GithubID(*payload.ID)).Only(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}

		p, err = h.ent.Pulls.Create().
			SetCreatedAt(*payload.CreatedAt).
			SetOwner(*payload.Head.Repo.Owner.Login).
			SetGithubID(*payload.ID).
			SetRepo(*payload.Head.Repo.Name).
			SetCreator(u).Save(ctx)
		if err != nil {
			return err
		}
	}

	if u.BangumiID == 0 && p.Comment == nil {
		c, _, err := h.github.PullRequests.CreateComment(ctx, p.Owner, p.Repo, int(p.GithubID), &github.PullRequestComment{
			Body: lo.ToPtr("请关联您的 bangumi ID 以方便进行贡献者统计\n\nhttps://contributors.bgm38.com/"),
		})

		if err != nil {
			return err
		}

		err = h.ent.Pulls.UpdateOne(p).SetComment(*c.ID).Exec(ctx)
		if err != nil {
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
		OnConflict().UpdateNewValues().Exec(ctx)
	if err != nil {
		logger.Err(err).Msg("failed to save authorized user to db")
		return err
	}

	return nil
}
