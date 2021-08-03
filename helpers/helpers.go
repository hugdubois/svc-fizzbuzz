// Package helpers provides some tests helpers functions.
package helpers

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"

	"github.com/hugdubois/svc-fizzbuzz/store"
	"github.com/sirupsen/logrus"
)

// CaptureOutput captures log and standard outputs.
func CaptureOutput(f func()) string {
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
	logrus.SetOutput(writer)
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

// WARN uses it only with test I don't known if there are some races conditions
var members = make(map[string]map[string]int64)

// MockHits is mockup of hits bag (store.Hits)
type MockHits struct{ Key string }

func NewMockHits(k string) store.Hitable {
	return MockHits{Key: k}
}

// Add incr the value of the member k in the hits bag
func (h MockHits) Add(k string, i int64) {
	if _, ok := members[h.Key]; !ok {
		members[h.Key] = make(map[string]int64)
	}

	if val, ok := members[h.Key][k]; ok {
		members[h.Key][k] = val + i
	} else {
		members[h.Key][k] = i
	}
}

// Top returns the most popular hit of hits bag
func (h MockHits) Top() (string, int64, error) {
	current, count := "", int64(0)
	if hits, ok := members[h.Key]; ok {
		for member, hit := range hits {
			if hit > count {
				current = member
				count = hit
			}
		}
	}

	return current, count, nil
}

// Reset the hits bag
func (h MockHits) Reset() {
	if _, ok := members[h.Key]; ok {
		members[h.Key] = make(map[string]int64)
	}
}
