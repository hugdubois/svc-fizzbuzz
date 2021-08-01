package cmd

import (
	"testing"
)

func TestVersionCmd(t *testing.T) {
	out := captureOutput(func() {
		rootCmd.SetArgs([]string{"version"})
		rootCmd.Execute()
	})
	expected := "svc-fizzbuzz version latest - svc-fizzbuzz-latest"
	if out != expected {
		t.Fatalf("svc-fizzbuzz version must return service version, got %s but want %s", out, expected)
	}
}
