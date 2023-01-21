package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/go-github/v49/github"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/samber/lo"

	"github-bot/ent"
	"github-bot/ent/user"
)

type PRHandle struct {
	logger zerolog.Logger
	ent    *ent.Client
	github *github.Client
}

func (h PRHandle) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	var payload github.PullRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		return err
	}

	h.logger.Debug().Interface("payload", payload).Msg("new event")

	h.handle(ctx, payload)

	return c.String(http.StatusOK, "Hello, World!")
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

	p, err := h.ent.Pulls.Create().
		SetCreatedAt(*payload.CreatedAt).
		SetOwner(*payload.Head.Repo.Owner.Login).
		SetGithubID(*payload.ID).
		SetRepo(*payload.Head.Repo.Name).
		SetCreator(u).Save(ctx)
	if err != nil {
		return err
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
