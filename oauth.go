package main

import (
	"context"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-github/v81/github"
	"github.com/kataras/go-sessions/v3"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/trim21/errgo"
	"golang.org/x/oauth2"

	"github-bot/config"
	"github-bot/ent/pulls"
	"github-bot/ent/user"
)

func (h PRHandle) setupGithubOAuth(e *echo.Echo) {
	conf := &oauth2.Config{
		ClientID:     config.GitHubOAuthAppID,
		ClientSecret: config.GitHubOAuthSecret,
		RedirectURL:  "https://contributors.bgm38.com/oauth/github/callback",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://github.com/login/oauth/access_token",
			AuthURL:  "https://github.com/login/oauth/authorize",
		},
	}

	{
		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

		e.GET("/oauth/github", func(c echo.Context) error {
			return c.Redirect(http.StatusFound, url)
		})
	}

	e.GET("/oauth/github/callback", func(c echo.Context) error {
		code := c.QueryParams().Get("code")
		if code == "" {
			return c.String(http.StatusBadRequest, "missing code")
		}
		token, err := conf.Exchange(c.Request().Context(), code)
		if err != nil {
			log.Err(err).Msg("failed to auth")
			return errgo.Trace(err)
		}

		gh := github.NewClient(oauth2.NewClient(context.TODO(), oauth2.StaticTokenSource(token)))

		u, _, err := gh.Users.Get(c.Request().Context(), "")
		if err != nil {
			log.Err(err).Msg("failed to get github user info")
			return errgo.Trace(err)
		}

		s := session.Start(c.Response(), c.Request())
		s.Set("github_id", int64(u.GetID()))

		return c.Redirect(http.StatusFound, "/")
	})
}

func (h PRHandle) setupBangumiOAuth(e *echo.Echo) {
	conf := &oauth2.Config{
		ClientID:     config.BangumiClientID,
		ClientSecret: config.BangumiClientSecret,
		RedirectURL:  "https://contributors.bgm38.com/oauth/bangumi/callback",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://bgm.tv/oauth/access_token",
			AuthURL:  "https://bgm.tv/oauth/authorize",
		},
	}

	{
		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

		e.GET("/oauth/bangumi", func(c echo.Context) error {
			return c.Redirect(http.StatusFound, url)
		})
	}

	var client = resty.New().SetHeader("user-agent", "bangumi-github-bot")

	e.GET("/oauth/bangumi/callback", func(c echo.Context) error {
		ctx := c.Request().Context()
		code := c.QueryParams().Get("code")
		if code == "" {
			return c.String(http.StatusBadRequest, "missing code")
		}

		token, err := conf.Exchange(c.Request().Context(), code)
		if err != nil {
			log.Err(err).Msg("failed to auth")
			return errgo.Trace(err)
		}

		var data struct {
			ID int `json:"id"`
		}

		res, err := client.R().SetHeader(echo.HeaderAuthorization, "Bearer "+token.AccessToken).SetResult(&data).Get("https://api.bgm.tv/v0/me")
		if err != nil {
			log.Err(err).Msg("failed to fetch user info from API")
			return errgo.Trace(err)
		}

		if res.StatusCode() > 300 {
			log.Error().
				Int("response_code", res.StatusCode()).
				Str("response_body", res.String()).
				Msg("failed to fetch user info, wrong http code")
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"code": res.StatusCode(),
				"body": res.String(),
			})
		}

		s := session.Start(c.Response(), c.Request())
		s.Set("bangumi_id", int64(data.ID))

		if err := h.afterOauth(ctx, s); err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/")
	})
}

func (h PRHandle) afterOauth(ctx context.Context, s *sessions.Session) error {
	githubId := s.GetInt64Default("github_id", 0)
	bangumiId := s.GetInt64Default("bangumi_id", 0)

	if githubId == 0 || bangumiId == 0 {
		return nil
	}

	err := h.ent.User.Create().SetGithubID(githubId).SetBangumiID(bangumiId).
		OnConflict(sql.ConflictColumns(user.FieldGithubID)).UpdateBangumiID().Exec(ctx)
	if err != nil {
		log.Err(err).Msg("failed to save authorized user to db")
		return errgo.Trace(err)
	}

	{
		prs, err := h.ent.Pulls.Query().Where(
			pulls.HasCreatorWith(user.GithubID(githubId)),
			pulls.CheckRunResult(checkRunActionRequired),
			pulls.CheckRunIDNEQ(0),
			pulls.MergedAtIsNil(),
		).All(ctx)
		if err != nil {
			log.Err(err).Msg("failed to get pulls")
			return errgo.Trace(err)
		}

		for _, pr := range prs {
			_, _, err := h.g.Checks.UpdateCheckRun(ctx, pr.Owner, pr.Repo, pr.CheckRunID, github.UpdateCheckRunOptions{
				Name:       githubCheckRunName,
				Conclusion: lo.ToPtr(checkRunSuccess),
				Output: &github.CheckRunOutput{
					Title:   lo.ToPtr(""),
					Summary: &successMessage,
				},
			})

			if err != nil {
				return errgo.Trace(err)
			}

			if err := h.ent.Pulls.UpdateOne(pr).
				SetCheckRunResult(checkRunSuccess).
				SetCheckRunID(0).
				Exec(ctx); err != nil {
				return errgo.Trace(err)
			}
		}
	}

	{
		prs, err := h.ent.Pulls.Query().Where(
			pulls.HasCreatorWith(user.GithubID(githubId)),
			pulls.CommentNEQ(0),
			pulls.MergedAtIsNil(),
		).All(ctx)
		if err != nil {
			log.Err(err).Msg("failed to get pulls")
			return errgo.Trace(err)
		}

		for _, pr := range prs {
			if _, err := h.g.Issues.DeleteComment(ctx, pr.Owner, pr.Repo, pr.Comment); err != nil {
				return errgo.Trace(err)
			}

			if err := h.ent.Pulls.UpdateOne(pr).SetComment(0).Exec(ctx); err != nil {
				return errgo.Trace(err)
			}
		}
	}

	return nil
}
