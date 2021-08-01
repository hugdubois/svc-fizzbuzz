package cmd

import (
	"regexp"
	"testing"

	"github.com/hugdubois/svc-fizzbuzz/helpers"
)

func Test_RootCmd(t *testing.T) {
	out := helpers.CaptureOutput(func() {
		Execute()
	})
	matched, err := regexp.MatchString(`Usage:`, out)
	if matched != true || err != nil {
		t.Fatalf("Command must display usage got:\n%s", out)
	}
	matched, err = regexp.MatchString(`svc-fizzbuzz \[command\]`, out)
	if matched != true || err != nil {
		t.Fatalf("Command must display usage got:\n%s", out)
	}
}
