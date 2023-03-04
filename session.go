package main

import (
	"github.com/jaevor/go-nanoid"
	"github.com/kataras/go-sessions/v3"
	"github.com/samber/lo"
)

var session = sessions.New(sessions.Config{
	SessionIDGenerator:          lo.Must(nanoid.Standard(32)),
	DisableSubdomainPersistence: true,
})
