package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hugdubois/svc-fizzbuzz/service/handlers"
)

func Test_NewService(t *testing.T) {
	svc := NewService()
	if got, want := svc.Name, name; got != want {
		t.Fatalf("Bad service name, got %s but want %s", got, want)
	}
	if got, want := svc.Version, version; got != want {
		t.Fatalf("Bad service version, got %s but want %s", got, want)
	}
}

func Test_StatusHandler(t *testing.T) {
	var statusMsg StatusResponse

	svc := NewService()
	ts := httptest.NewServer(svc.NewRouter())
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/status", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
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
	if got, want := statusMsg.Alive, true; got != want {
		t.Fatalf("Wrong version return, got %t but want %t", got, want)
	}
}

func Test_VersionHandler(t *testing.T) {
	svc := NewService()
	ts := httptest.NewServer(svc.NewRouter())
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/version", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}
	var versionMsg Service
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
}

func Test_Index(t *testing.T) {
	var versionMsg Service

	svc := NewService()
	ts := httptest.NewServer(svc.NewRouter())
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/", nil)
	resp, _ := http.DefaultClient.Do(req)
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
}

func Test_NotFound(t *testing.T) {
	var notFoundMsg handlers.ErrorMessage

	svc := NewService()
	ts := httptest.NewServer(svc.NewRouter())
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/not_found", nil)
	resp, _ := http.DefaultClient.Do(req)
	if got, want := resp.StatusCode, 404; got != want {
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
	if got, want := notFoundMsg.Code, 404; got != want {
		t.Fatalf("Wrong code return, got %d but want %d", got, want)
	}
	if got, want := notFoundMsg.Message, "Not Found"; got != want {
		t.Fatalf("Wrong message return, got %s but want %s", got, want)
	}
}
