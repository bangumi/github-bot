package main

import (
	"context"
	"encoding/base64"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/rs/zerolog"
	"github.com/samber/lo"

	"github-bot/ent"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

var version = "development"

var installationID = lo.Must(strconv.ParseInt(os.Getenv("GITHUB_BANGUMI_INSTALLATION_ID"), 10, 64))

func main() {
	client, err := ent.Open("postgres", os.Getenv("PG_OPTIONS"))
	if err != nil {
		logger.Fatal().Err(err).Msg("failed opening connection to pg")
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal().Err(err).Msg("failed creating schema resources")
	}

	// Echo instance
	e := echo.New()
	e.HideBanner = true

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				logger.Err(err).Msg("internal error")
			}

			return err
		}
	})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("x-version", version)
			return next(c)
		}
	})
	// Middleware
	e.Use(middleware.Recover())

	h := PRHandle{
		ent: client,
		app: getGithubAppClient(),
	}

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

func getGithubAppClient() githubapp.ClientCreator {
	pemRaw := lo.Must(base64.StdEncoding.DecodeString(os.Getenv("GITHUB_APP_CERT_PRIVATE")))

	return lo.Must(githubapp.NewDefaultCachingClientCreator(githubapp.Config{
		V3APIURL: "https://api.github.com/",
		App: struct {
			IntegrationID int64  `yaml:"integration_id" json:"integrationId"`
			WebhookSecret string `yaml:"webhook_secret" json:"webhookSecret"`
			PrivateKey    string `yaml:"private_key" json:"privateKey"`
		}{
			IntegrationID: 289933,
			PrivateKey:    string(pemRaw),
		},
	}))
}
