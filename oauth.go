package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"

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

	bangumiId := getBangumiID(sess)
	githubId := getGithubID(sess)

	var html string

	if bangumiId == 0 {
		html += `<p> bangumi 未链接，请认证 <a href="/oauth/bangumi">bangumi oauth</a> </p>`
	} else {
		html += fmt.Sprintf(`<p> bangumi id %d </p>`, bangumiId)
	}

	if githubId == 0 {
		html += `<p> github 未链接，请认证 <a href="/oauth/github">github oauth</a> </p>`
	} else {
		html += fmt.Sprintf(`<p> github id %d </p>`, githubId)
	}

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

func setupBangumiOAuth(e *echo.Echo) {
	ctx := context.Background()

	conf := &oauth2.Config{
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://provider.com/o/oauth2/token",
			AuthURL:  "https://provider.com/o/oauth2/auth",
		},
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the custom HTTP client when requesting a token.
	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	e.GET("/oauth/bangumi", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "")
	})

	e.GET("/oauth/bangumi/callback", func(c echo.Context) error {
		token, err := conf.Exchange(ctx, c.QueryParams().Get("code"))
		if err != nil {
			logger.Err(err).Msg("failed to auth")
			return err
		}

		_ = token

		s, _ := session.Get("session", c)
		if s == nil {
			return nil
		}
		s.Options = &sessions.Options{Path: "/", HttpOnly: true}

		s.Values["bangumi_id"] = 1

		err = s.Save(c.Request(), c.Response())
		if err != nil {
			return err
		}

		return nil
	})
}

func githubHandler(c echo.Context) error {
	s, _ := session.Get("session", c)
	if s == nil {
		return nil
	}
	s.Options = &sessions.Options{Path: "/", HttpOnly: true}

	s.Values["github_id"] = 1

	err := s.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}

	return nil
}
