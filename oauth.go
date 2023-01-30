package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"

	"github.com/labstack/echo/v4"
)

func (h PRHandle) setupGithubOAuth(e *echo.Echo) {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_OAUTH_APP_ID"),
		ClientSecret: os.Getenv("GITHUB_OAUTH_APP_SECRET"),
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
			logger.Err(err).Msg("failed to auth")
			return err
		}

		gh := github.NewClient(oauth2.NewClient(context.TODO(), oauth2.StaticTokenSource(token)))

		u, _, err := gh.Users.Get(c.Request().Context(), "")
		if err != nil {
			logger.Err(err).Msg("failed to get github user info")
			return err
		}

		s := session.Start(c.Response(), c.Request())
		s.Set("github_id", int(*u.ID))

		return c.Redirect(http.StatusFound, "/")
	})
}

func (h PRHandle) setupBangumiOAuth(e *echo.Echo) {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("BANGUMI_OAUTH_APP_ID"),
		ClientSecret: os.Getenv("BANGUMI_OAUTH_APP_SECRET"),
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
		code := c.QueryParams().Get("code")
		if code == "" {
			return c.String(http.StatusBadRequest, "missing code")
		}

		token, err := conf.Exchange(c.Request().Context(), code)
		if err != nil {
			logger.Err(err).Msg("failed to auth")
			return err
		}

		var data struct {
			ID int `json:"id"`
		}

		res, err := client.R().SetHeader(echo.HeaderAuthorization, "Bearer "+token.AccessToken).SetResult(&data).Get("https://api.bgm.tv/v0/me")
		if err != nil {
			logger.Err(err).Msg("failed to fetch user info from API")
			return err
		}

		if res.StatusCode() > 300 {
			logger.Error().
				Int("response_code", res.StatusCode()).
				Str("response_body", res.String()).
				Msg("failed to fetch user info, wrong http code")
			return c.NoContent(http.StatusInternalServerError)
		}

		s := session.Start(c.Response(), c.Request())
		s.Set("bangumi_id", data.ID)

		return c.Redirect(http.StatusFound, "/")
	})
}
