package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// assertCorsMiddleware is a helper function to test the valid the
// CorsMiddleware.
func assertCorsMiddleware(t *testing.T, verb string) {
	t.Helper()

	ts := httptest.NewServer(NewCors("*")(http.DefaultServeMux))
	defer ts.Close()

	req, _ := http.NewRequest(verb, ts.URL+"/cors", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("got %#v, want %#v", got, want)
	}

	if v, ok := resp.Header["Access-Control-Allow-Origin"]; !ok || v[0] != "*" {
		t.Fatalf("expected origin CORS header: got %s", resp.Header)
	}
	if v, ok := resp.Header["Access-Control-Allow-Methods"]; !ok || v[0] != "GET,HEAD,OPTIONS" {
		t.Fatalf("expected origin CORS header: got %s", resp.Header)
	}
}

// assertCorsMiddlewareEmpty is a helper function to test the valid the
// CorsMiddleware (when origin is not set).
func assertCorsMiddlewareEmpty(t *testing.T, verb string) {
	t.Helper()

	ts := httptest.NewServer(NewCors("")(http.DefaultServeMux))
	defer ts.Close()

	req, _ := http.NewRequest(verb, ts.URL+"/cors-empty", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("got %#v, want %#v", got, want)
	}

	if _, ok := resp.Header["Access-Control-Allow-Origin"]; ok {
		t.Fatalf("not expected origin CORS header: got %s", resp.Header)
	}

	if _, ok := resp.Header["Access-Control-Allow-Methods"]; ok {
		t.Fatalf("not expected origin CORS header: got %s", resp.Header)
	}
}

// Test_CorsMiddleware provides the CorsMiddleware test.
func Test_CorsMiddleware(t *testing.T) {
	http.HandleFunc("/cors", func(resp http.ResponseWriter, req *http.Request) {})
	assertCorsMiddleware(t, http.MethodOptions)
	assertCorsMiddleware(t, http.MethodHead)
	assertCorsMiddleware(t, http.MethodGet)
}

// Test_CorsMiddleware provides the CorsMiddleware test (when origin is not set).
func Test_CorsMiddlewareWithEmptyOrigin(t *testing.T) {
	http.HandleFunc("/cors-empty", func(resp http.ResponseWriter, req *http.Request) {})
	assertCorsMiddlewareEmpty(t, http.MethodOptions)
	assertCorsMiddlewareEmpty(t, http.MethodHead)
	assertCorsMiddlewareEmpty(t, http.MethodGet)
}
