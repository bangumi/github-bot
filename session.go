package main

import (
	"github.com/kataras/go-sessions/v3"
)

var session = sessions.New(sessions.Config{
	Cookie:                      "jwt-session",
	DisableSubdomainPersistence: true,
})
