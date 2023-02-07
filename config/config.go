package config

import (
	"encoding/base64"
	"os"
	"strconv"

	"github.com/samber/lo"
)

var WebhookSecret = []byte(os.Getenv("GITHUB_APP_WEBHOOK_SECRET"))
var InstallationID = lo.Must(strconv.ParseInt(os.Getenv("GITHUB_BANGUMI_INSTALLATION_ID"), 10, 64))
var PgOption = os.Getenv("PG_OPTIONS")

var GitHubAppPrivateKey = string(lo.Must(base64.StdEncoding.DecodeString(os.Getenv("GITHUB_APP_CERT_PRIVATE"))))

var GitHubOAuthAppID = os.Getenv("GITHUB_OAUTH_APP_ID")
var GitHubOAuthSecret = os.Getenv("GITHUB_OAUTH_APP_SECRET")

var BangumiClientID = os.Getenv("BANGUMI_OAUTH_APP_ID")
var BangumiClientSecret = os.Getenv("BANGUMI_OAUTH_APP_SECRET")

var HTTPPost = os.Getenv("HTTP_PORT")

var HTTPHost = os.Getenv("HTTP_HOST")
