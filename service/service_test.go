package service

import (
	"testing"
)

func init() {
	pingStore = func() (string, error) {
		return "PONG", nil
	}
}

// Test_NewService provides the NewService test.
func Test_NewService(t *testing.T) {
	svc := NewService()
	if got, want := svc.Name, name; got != want {
		t.Fatalf("Bad service name, got %s but want %s", got, want)
	}
	if got, want := svc.Version, version; got != want {
		t.Fatalf("Bad service version, got %s but want %s", got, want)
	}
}
