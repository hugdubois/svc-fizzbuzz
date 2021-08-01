package cmd

import (
	"regexp"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestServeCmd(t *testing.T) {
	srv := getServer()
	if got, want := srv.Addr, defautAddress; got != want {
		t.Fatalf("Server addres flag error, got %s but want %s", got, want)
	}

	currentDebugMode := debugMode
	if debugMode == false {
		debugMode = true
	}
	initLogger()
	out := captureOutput(func() {
		log.Debug("must appear in log")
	})

	matched, err := regexp.MatchString(`must appear in log`, out)
	if matched != true || err != nil {
		t.Fatalf("debug mode is not active got:\n%s", out)
	}

	debugMode = currentDebugMode
}
