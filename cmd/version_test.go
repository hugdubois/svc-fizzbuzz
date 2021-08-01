package cmd

import (
	"testing"

	"github.com/hugdubois/svc-fizzbuzz/helpers"
)

func Test_VersionCmd(t *testing.T) {
	out := helpers.CaptureOutput(func() {
		rootCmd.SetArgs([]string{"version"})
		rootCmd.Execute()
	})
	expected := "svc-fizzbuzz version latest - svc-fizzbuzz-latest"
	if out != expected {
		t.Fatalf("svc-fizzbuzz version must return service version, got %s but want %s", out, expected)
	}
}
