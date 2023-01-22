package main

import (
	"context"
	"os"

	"github.com/google/go-github/v49/github"
	"github.com/kataras/go-sessions/v3/sessiondb/boltdb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"

	"github-bot/ent"
)

var logger zerolog.Logger

func main() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &logger

	client, err := ent.Open("sqlite3", "file:./data/ent.sqlite?cache=shared&_fk=1")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed opening connection to sqlite")
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal().Err(err).Msg("failed creating schema resources")
	}

	b, err := boltdb.New("./session.bolt", 0644)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to open bolt file to store session")
	}
	session.UseDatabase(b)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_COMMITTER_ACCESS_TOKEN")},
	)

	h := PRHandle{logger: logger, ent: client, github: github.NewClient(oauth2.NewClient(ctx, ts))}

	// Routes
	e.POST("/event-pr", h.Handle)
	e.GET("/", h.Index)
	h.setupGithubOAuth(e)
	h.setupBangumiOAuth(e)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8090"
	}

	host := os.Getenv("HTTP_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	// Start server
	e.Logger.Fatal(e.Start(host + ":" + port))
}
