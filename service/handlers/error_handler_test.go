package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ErrorHandler(t *testing.T) {
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		ErrorHandler(w, r, http.StatusBadRequest, "Bad Parameter")
	})
	testserver := httptest.NewServer(http.DefaultServeMux)

	req, _ := http.NewRequest("GET", testserver.URL+"/error", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, http.StatusBadRequest; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	var errorMsg ErrorMessage
	err = json.Unmarshal(data, &errorMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}
	if got, want := errorMsg.Code, http.StatusBadRequest; got != want {
		t.Fatalf("Wrong code return, got %d but want %d", got, want)
	}
	if got, want := errorMsg.Message, "Bad Parameter"; got != want {
		t.Fatalf("Wrong message return, got %s but want %s", got, want)
	}
}
