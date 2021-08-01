package middlewares

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// A middleware constructor
func taggedMiddleware(tag string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(tag))
			h.ServeHTTP(w, r)
		})
	}
}

func Test_UseMiddleware(t *testing.T) {
	useMiddleware := UseMiddleware(
		taggedMiddleware("A"),
		taggedMiddleware("B"),
		taggedMiddleware("C"),
	)

	useEmptyMiddleware := UseMiddleware()

	http.Handle(
		"/endpoint",
		useMiddleware(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "endpoint")
		}),
	)

	http.Handle(
		"/empty-middleware-endpoint",
		useEmptyMiddleware(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "empty middleware endpoint")
		}),
	)

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/endpoint", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}
	if got, want := string(data[:]), "ABCendpoint"; got != want {
		t.Fatalf("Invalid body, got %s but want %s", got, want)
	}

	req, _ = http.NewRequest("GET", ts.URL+"/empty-middleware-endpoint", nil)
	resp, _ = http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}
	if got, want := string(data[:]), "empty middleware endpoint"; got != want {
		t.Fatalf("Invalid body, got %s but want %s", got, want)
	}
}
