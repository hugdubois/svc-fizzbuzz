package cmd

import (
	"regexp"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/hugdubois/svc-fizzbuzz/helpers"
)

// Test_ServeCmd provides the ServeCmd test.
func Test_ServeCmd(t *testing.T) {
	srv := getServer()

	if got, want := srv.Addr, DefautAddress; got != want {
		t.Fatalf("Server address flag error, got %s but want %s", got, want)
	}

	if got, want := srv.ReadTimeout, DefaultReadTimeout; got != want {
		t.Fatalf("Server read-timeout flag error, got %s but want %s", got, want)
	}

	if got, want := srv.WriteTimeout, DefaultWriteTimeout; got != want {
		t.Fatalf("Server write-timeout flag error, got %s but want %s", got, want)
	}

	currentDebugMode := debugMode
	if debugMode == false {
		debugMode = true
	}

	out := helpers.CaptureOutput(func() {
		initLogger()
		log.Print("must appear in log")
	})

	matched, err := regexp.MatchString(`must appear in log`, out)
	if matched != true || err != nil {
		t.Fatalf("debug mode is not active got:\n%s", out)
	}

	debugMode = currentDebugMode
}
