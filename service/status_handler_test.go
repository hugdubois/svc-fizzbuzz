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

// Test_StatusHandler provides the StatusHandler test.
func Test_StatusHandler(t *testing.T) {
	var (
		statusMsg StatusResponse
		resp      *http.Response
	)

	svc := NewService()
	ts := httptest.NewServer(svc.NewRouter("*"))
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/status", nil)

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

	err = json.Unmarshal(data, &statusMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}
	if got, want := statusMsg.SvcAlive, true; got != want {
		t.Fatalf("Wrong status, service must be alive, got %t but want %t", got, want)
	}
	if got, want := statusMsg.StoreAlive, true; got != want {
		t.Fatalf("Wrong status, store must be reachable, got %t but want %t", got, want)
	}

	matched, err := regexp.MatchString(`uri=/status `, outputLog)
	if matched != true || err != nil {
		t.Fatalf("request is not logged :\n%s", outputLog)
	}
}
