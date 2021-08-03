package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Test_MetricsMiddleware provides MetricsMiddleware test.
func Test_MetricsMiddleware(t *testing.T) {
	http.HandleFunc("/a-endpoint", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	http.Handle("/metrics", promhttp.Handler())

	ts := httptest.NewServer(NewMetrics()(http.DefaultServeMux))
	defer ts.Close()

	// call metrics endpoint before /a-endpoint so it must not be found
	req, _ := http.NewRequest("GET", ts.URL+"/metrics", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}
	matched, err := regexp.Match(`/a-endpoint`, data)
	if matched != false || err != nil {
		t.Fatalf("Invalid regexp match %t %v", matched, err)
	}

	// call a monitored endpoint
	req, _ = http.NewRequest("GET", ts.URL+"/a-endpoint", nil)
	resp, _ = http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}
	if got, want := string(data[:]), "Hello world"; got != want {
		t.Fatalf("Wrong version return, got %s but want %s", got, want)
	}

	// call metrics endpoint /a-endpoint must be found
	req, _ = http.NewRequest("GET", ts.URL+"/metrics", nil)
	resp, _ = http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}
	matched, err = regexp.Match(`/a-endpoint`, data)
	if matched == false || err != nil {
		t.Fatalf("Invalid regexp not match %t %v", matched, err)
	}
}
