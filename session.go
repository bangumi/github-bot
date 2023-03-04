package main

import (
	"github.com/kataras/go-sessions/v3"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var session = sessions.New(sessions.Config{
	SessionIDGenerator: func() string {
		return gonanoid.Must()
	},
	DisableSubdomainPersistence: true,
})
