package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"

	"github-bot/config"
	"github-bot/ent"
)

var version = "development"

func main() {
	client, err := ent.Open("postgres", config.PgOption)
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to pg")
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	// Echo instance
	e := echo.New()
	e.HideBanner = true

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				if _, ok := err.(*echo.HTTPError); !ok {
					log.Err(err).Msg("internal error")
				}
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

	appClient := getGithubAppClient()
	h := PRHandle{
		ent: client,
		app: appClient,
		g:   lo.Must(appClient.NewInstallationClient(config.InstallationID)),
	}

	// Routes
	e.POST("/event-pr", h.Handle)
	e.GET("/", h.Index)
	h.setupGithubOAuth(e)
	h.setupBangumiOAuth(e)

	port := config.HTTPPost
	if port == "" {
		port = "8090"
	}

	host := config.HTTPHost
	if host == "" {
		host = "127.0.0.1"
	}

	// Start server
	log.Fatal().Err(e.Start(host + ":" + port)).Msg("failed to start server")
}

func getGithubAppClient() githubapp.ClientCreator {
	return lo.Must(githubapp.NewDefaultCachingClientCreator(githubapp.Config{
		V3APIURL: "https://api.github.com/",
		App: struct {
			IntegrationID int64  `yaml:"integration_id" json:"integrationId"`
			WebhookSecret string `yaml:"webhook_secret" json:"webhookSecret"`
			PrivateKey    string `yaml:"private_key" json:"privateKey"`
		}{
			IntegrationID: 289933,
			PrivateKey:    config.GitHubAppPrivateKey,
		},
	}))
}
