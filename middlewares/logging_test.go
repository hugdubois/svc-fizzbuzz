package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/hugdubois/svc-fizzbuzz/helpers"
)

// Test_LoggingMiddleware provides the LoggingMiddleware test.
func Test_LoggingMiddleware(t *testing.T) {
	http.HandleFunc("/logged-endpoint", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "response body")
	})

	ts := httptest.NewServer(NewLogging("svc-fizzbuzz")(http.DefaultServeMux))
	defer ts.Close()

	var resp *http.Response

	req, _ := http.NewRequest("GET", ts.URL+"/logged-endpoint", nil)

	outputLog := helpers.CaptureOutput(func() {
		resp, _ = http.DefaultClient.Do(req)
	})

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	if got, want := string(data[:]), "response body"; got != want {
		t.Fatalf("Wrong body return, got %s but want %s", got, want)
	}

	matched, err := regexp.MatchString(`uri=/logged-endpoint `, outputLog)
	if matched != true || err != nil {
		t.Fatalf("request is not logged :\n%s", outputLog)
	}

	// some custom headers must appear in log
	req.Header.Add("X-Forwarded-For", "192.168.0.13")
	req.Header.Add("X-Request-Id", "13")

	outputLog = helpers.CaptureOutput(func() {
		resp, _ = http.DefaultClient.Do(req)
	})

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	if got, want := string(data[:]), "response body"; got != want {
		t.Fatalf("Wrong body return, got %s but want %s", got, want)
	}

	matched, err = regexp.MatchString(`uri=/logged-endpoint `, outputLog)
	if matched != true || err != nil {
		t.Fatalf("request is not logged :\n%s", outputLog)
	}

	matched, err = regexp.MatchString(`request_id=13 `, outputLog)
	if matched != true || err != nil {
		t.Fatalf("X-Request-Id header is not logged :\n%s", outputLog)
	}

	matched, err = regexp.MatchString(`forwarded_for=192.168.0.13 `, outputLog)
	if matched != true || err != nil {
		t.Fatalf("X-Forwarded-For header is not logged :\n%s", outputLog)
	}
}
