package main

import (
	"context"
	"net/http"
	"os"

	"github.com/google/go-github/v49/github"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"

	"github-bot/ent"
)

var logger zerolog.Logger

func main() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &logger

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed opening connection to sqlite")
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal().Err(err).Msg("failed creating schema resources")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tr := &AddHeaderTransport{T: http.DefaultTransport, token: os.Getenv("GITHUB_COMMITTER_ACCESS_TOKEN")}

	h := PRHandle{logger: logger, ent: client, github: github.NewClient(&http.Client{Transport: tr})}

	// Routes
	e.POST("/event-pr", h.Handle)
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET("/", Index)
	setupBangumiOAuth(e)
	setupGithubOAuth(e)

	// Start server
	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}

type AddHeaderTransport struct {
	T     http.RoundTripper
	token string
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+os.Getenv("GITHUB_COMMITTER_ACCESS_TOKEN"))
	return adt.T.RoundTrip(req)
}
