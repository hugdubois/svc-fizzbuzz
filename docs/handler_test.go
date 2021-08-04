package docs

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

// Test_Handler provides the Handler test.
func Test_VersionHandler(t *testing.T) {

	http.Handle("/swagger.json", Handler())
	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/swagger.json", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	matched, err := regexp.Match(`"swagger": "2.0"`, data)
	if matched != true || err != nil {
		t.Fatalf("Invalid swagger definition :\n%s", string(data[:]))
	}
}
