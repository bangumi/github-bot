package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v49/github"
	"golang.org/x/oauth2"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	sess, _ := session.Get("session", c)
	if sess == nil {
		return nil
	}
	sess.Options = &sessions.Options{Path: "/", HttpOnly: true}

	githubId := getGithubID(sess)

	var html string

	if githubId == 0 {
		return c.HTML(http.StatusOK, `<p> github 未链接，请认证 <a href="/oauth/github">github oauth</a> </p>`)
	}

	html += fmt.Sprintf(`<p> github id %d </p>`, githubId)

	bangumiId := getBangumiID(sess)
	if bangumiId == 0 {
		return c.HTML(http.StatusOK, `<p> bangumi 未链接，请认证 <a href="/oauth/bangumi">bangumi oauth</a> </p>`)
	}

	html += fmt.Sprintf(`<p> bangumi id %d </p>`, bangumiId)

	return c.HTML(http.StatusOK, html)
}

func getBangumiID(s *sessions.Session) int {
	raw, ok := s.Values["bangumi_id"]
	if !ok {
		return 0
	}
	return raw.(int)
}

func getGithubID(s *sessions.Session) int {
	raw, ok := s.Values["github_id"]
	if !ok {
		return 0
	}
	return raw.(int)
}

var client = resty.New()

func setupBangumiOAuth(e *echo.Echo) {
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
		fmt.Printf("Visit the URL for the auth dialog: %v", url)

		e.GET("/oauth/bangumi", func(c echo.Context) error {
			return c.Redirect(http.StatusFound, url)
		})
	}

	e.GET("/oauth/bangumi/callback", func(c echo.Context) error {
		token, err := conf.Exchange(c.Request().Context(), c.QueryParams().Get("code"))
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

		s, _ := session.Get("session", c)
		if s == nil {
			return nil
		}
		s.Options = &sessions.Options{Path: "/", HttpOnly: true}

		s.Values["bangumi_id"] = data.ID

		err = s.Save(c.Request(), c.Response())
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, "/")
	})
}

func setupGithubOAuth(e *echo.Echo) {
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
		fmt.Printf("Visit the URL for the auth dialog: %v", url)

		e.GET("/oauth/github", func(c echo.Context) error {
			return c.Redirect(http.StatusFound, url)
		})
	}

	e.GET("/oauth/github/callback", func(c echo.Context) error {
		token, err := conf.Exchange(c.Request().Context(), c.QueryParams().Get("code"))
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

		s, _ := session.Get("session", c)
		if s == nil {
			return nil
		}
		s.Options = &sessions.Options{Path: "/", HttpOnly: true}

		s.Values["github_id"] = int(*u.ID)

		err = s.Save(c.Request(), c.Response())
		if err != nil {
			return err
		}

		return nil
	})
}
