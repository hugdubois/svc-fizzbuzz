package hits

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

var (
	// hitsTests store the fizzbuzz hits
	hitsTests *Hits
)

func init() {
	hitsTests = NewHits("hitsTests")
}

func Test_HitsHandler(t *testing.T) {
	hitsTests.Init()
	http.HandleFunc("/hits", Handler)
	http.HandleFunc("/endpoint-hited-1", func(w http.ResponseWriter, r *http.Request) {
		hitsTests.Add("endpoint-hited-1", 1)
	})
	http.HandleFunc("/endpoint-hited-2", func(w http.ResponseWriter, r *http.Request) {
		hitsTests.Add("endpoint-hited-2", 1)
	})

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	// first call to hits
	req, _ := http.NewRequest("GET", ts.URL+"/hits", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	matched, err := regexp.Match(`endpoint-hited-1`, data)
	if matched == true || err != nil {
		t.Fatalf("Invalid regexp match %t %v", matched, err)
	}

	// call a hited endpoint 2 first (to test coverage on sort)
	req, _ = http.NewRequest("GET", ts.URL+"/endpoint-hited-2", nil)
	resp, _ = http.DefaultClient.Do(req)
	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	// call a hited endpoint 1
	req, _ = http.NewRequest("GET", ts.URL+"/endpoint-hited-1", nil)
	resp, _ = http.DefaultClient.Do(req)
	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	// second call to hits
	req, _ = http.NewRequest("GET", ts.URL+"/hits", nil)
	resp, _ = http.DefaultClient.Do(req)
	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	matched, err = regexp.Match(`endpoint-hited-1`, data)
	if matched == false || err != nil {
		t.Fatalf("Invalid regexp not match %t %v", matched, err)
	}

	// call a hited endpoint 2 two time it must be first on top
	req, _ = http.NewRequest("GET", ts.URL+"/endpoint-hited-2", nil)
	resp, _ = http.DefaultClient.Do(req)
	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	// hits top must be endpoint-hited-2 with 2 hits
	top, count := hitsTests.Top()
	if want := "endpoint-hited-2"; top != want {
		t.Fatalf("Bad most hited endpoint, got %s but want %s", top, want)
	}
	if want := int64(2); count != want {
		t.Fatalf("Bad most hited endpoint count, got %d but want %d", count, want)
	}

	hitsTests.Init()
}
