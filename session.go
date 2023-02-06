package main

import (
	"bytes"
	"encoding/gob"

	"github.com/kataras/go-sessions/v3"
)

var session = sessions.New(sessions.Config{
	Cookie:                      "jwt-session",
	DisableSubdomainPersistence: true,
})

var _ sessions.Transcoder = gobTranscoder{}

type gobTranscoder struct {
}

func (g gobTranscoder) Marshal(i interface{}) ([]byte, error) {
	var b = bytes.NewBuffer(nil)
	err := gob.NewEncoder(b).Encode(i)

	return b.Bytes(), err
}

func (g gobTranscoder) Unmarshal(b []byte, i interface{}) error {
	return gob.NewDecoder(bytes.NewBuffer(b)).Decode(&i)
}
