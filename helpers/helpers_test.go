package helpers

import (
	"fmt"
	"log"
	"regexp"
	"testing"
)

func Test_CaptureOutput(t *testing.T) {
	expected := "hello world"
	out := CaptureOutput(func() {
		fmt.Print(expected)
	})
	if out != expected {
		t.Fatalf("fmt.Print must be captured, got %s but want %s", out, expected)
	}

	out = CaptureOutput(func() {
		log.Print(expected)
	})
	matched, err := regexp.MatchString(expected, out)
	if matched != true || err != nil {
		t.Fatalf("log.Print must be captured, got %s but want %s", out, expected)
	}
}
