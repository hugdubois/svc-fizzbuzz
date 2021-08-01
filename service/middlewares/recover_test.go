package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/hugdubois/svc-fizzbuzz/helpers"
)

func assertRecover(t *testing.T, url string, fn http.HandlerFunc) {
	t.Helper()

	http.HandleFunc(url, fn)

	ts := httptest.NewServer(RecoverMiddleware(http.DefaultServeMux))
	defer ts.Close()

	var (
		req  *http.Request
		resp *http.Response
	)

	outputLog := helpers.CaptureOutput(func() {
		req, _ = http.NewRequest("GET", ts.URL+url, nil)
		resp, _ = http.DefaultClient.Do(req)
	})

	if got, want := resp.StatusCode, http.StatusInternalServerError; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	matched, err := regexp.MatchString(`There was a panic with this message:`, outputLog)
	if matched != true || err != nil {
		t.Fatalf("panic is not logged :\n%s", outputLog)
	}
}

func Test_RecoverMiddleware(t *testing.T) {
	assertRecover(t, "/recover-string", func(http.ResponseWriter, *http.Request) { panic("foo") })
	assertRecover(t, "/recover-error", func(http.ResponseWriter, *http.Request) { panic(fmt.Errorf("foo %s", "bar")) })
	assertRecover(t, "/recover-unknown", func(http.ResponseWriter, *http.Request) { panic(13) })
}
