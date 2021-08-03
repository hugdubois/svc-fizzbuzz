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

// Test_VersionHandler provides the VersionHandler test.
func Test_VersionHandler(t *testing.T) {
	var (
		resp       *http.Response
		versionMsg Service
	)

	svc := NewService()
	ts := httptest.NewServer(svc.NewRouter("*"))
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/version", nil)

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

	matched, err := regexp.MatchString(`uri=/version `, outputLog)
	if matched != true || err != nil {
		t.Fatalf("request is not logged :\n%s", outputLog)
	}
}
