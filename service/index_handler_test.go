package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/hugdubois/svc-fizzbuzz/helpers"
)

// Test_IndexHandler provides the FizzbuzzTopHandler test.
func Test_IndexHandler(t *testing.T) {
	var (
		versionMsg Service
		resp       *http.Response
	)

	svc := NewService()

	ts := httptest.NewServer(svc.NewRouter("*"))
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/", nil)

	outputLog := helpers.CaptureOutput(func() {
		resp, _ = http.DefaultClient.Do(req)
	})

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	err = json.Unmarshal(data, &versionMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}
	if got, want := versionMsg.Version, svc.Version; got != want {
		t.Fatalf("Wrong version return, got %s but want %s", got, want)
	}
	if got, want := versionMsg.Name, svc.Name; got != want {
		t.Fatalf("Wrong version return, got %s but want %s", got, want)
	}

	matched, err := regexp.MatchString(`uri=/ `, outputLog)
	if matched != true || err != nil {
		t.Fatalf("request is not logged :\n%s", outputLog)
	}
}

// Test_NotFound provides the IndexHandler test (when the url doesn't match
// anything).
func Test_NotFound(t *testing.T) {
	var (
		notFoundMsg ErrorMessage
		resp        *http.Response
	)

	svc := NewService()
	ts := httptest.NewServer(svc.NewRouter("*"))
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/not_found", nil)

	outputLog := helpers.CaptureOutput(func() {
		resp, _ = http.DefaultClient.Do(req)
	})

	if got, want := resp.StatusCode, http.StatusNotFound; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}
	err = json.Unmarshal(data, &notFoundMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}
	if got, want := notFoundMsg.Code, http.StatusNotFound; got != want {
		t.Fatalf("Wrong code return, got %d but want %d", got, want)
	}
	if got, want := notFoundMsg.Message, "Not Found"; got != want {
		t.Fatalf("Wrong message return, got %s but want %s", got, want)
	}

	matched, err := regexp.MatchString(`uri=/not_found `, outputLog)
	if matched != true || err != nil {
		t.Fatalf("request is not logged :\n%s", outputLog)
	}
}
