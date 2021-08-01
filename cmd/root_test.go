package cmd

import (
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
	"testing"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestRootCmd(t *testing.T) {
	out := captureOutput(func() {
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
