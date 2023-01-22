package main

import (
	"context"
	"os"

	"github.com/google/go-github/v49/github"
	"github.com/kataras/go-sessions/v3/sessiondb/boltdb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2"

	"github-bot/ent"
)

var logger zerolog.Logger

var version = "development"

func main() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &logger

	client, err := ent.Open("postgres", os.Getenv("PG_OPTIONS"))
	if err != nil {
		logger.Fatal().Err(err).Msg("failed opening connection to pg")
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal().Err(err).Msg("failed creating schema resources")
	}

	b, err := boltdb.New("./data/session.bolt", 0644)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to open bolt file to store session")
	}
	session.UseDatabase(b)

	// Echo instance
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("x-version", version)
			return next(c)
		}
	})
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
