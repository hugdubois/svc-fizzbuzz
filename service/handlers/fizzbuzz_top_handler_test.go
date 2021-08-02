package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFizzbuzzTopHandler(t *testing.T) {
	var fizzbuzzTopMsg FizzBuzzTopResponse
	fizzbuzzHits.Init()
	http.HandleFunc("/fizzbuzz/top", FizzBuzzTopHandler)

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	// first call to fizzbuzz/top
	req, _ := http.NewRequest("GET", ts.URL+"/fizzbuzz/top", nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	err = json.Unmarshal(data, &fizzbuzzTopMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}
	if got, want := fizzbuzzTopMsg.CountRequest, int64(0); got != want {
		t.Fatalf("Wrong version return, got %d but want %d", got, want)
	}

	// call a fizzbuzz with an identifiable string
	req, _ = http.NewRequest("GET", ts.URL+"/fizzbuzz?str1=ONE_CALL", nil)
	resp, _ = http.DefaultClient.Do(req)
	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	// second call to fizzbuzz/top
	req, _ = http.NewRequest("GET", ts.URL+"/fizzbuzz/top", nil)
	resp, _ = http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	err = json.Unmarshal(data, &fizzbuzzTopMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}
	if got, want := fizzbuzzTopMsg.CountRequest, int64(1); got != want {
		t.Fatalf("Wrong version return, got %d but want %d", got, want)
	}
	if got, want := fizzbuzzTopMsg.FizzBuzzParams.Str1, "ONE_CALL"; got != want {
		t.Fatalf("Wrong version return, got %s but want %s", got, want)
	}

	// call a fizzbuzz with an identifiable string
	req, _ = http.NewRequest("GET", ts.URL+"/fizzbuzz?str1=TWO_CALL", nil)
	resp, _ = http.DefaultClient.Do(req)
	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	// call a fizzbuzz with an identifiable string
	req, _ = http.NewRequest("GET", ts.URL+"/fizzbuzz?str1=TWO_CALL", nil)
	resp, _ = http.DefaultClient.Do(req)
	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	// third call to fizzbuzz/top
	req, _ = http.NewRequest("GET", ts.URL+"/fizzbuzz/top", nil)
	resp, _ = http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, 200; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	err = json.Unmarshal(data, &fizzbuzzTopMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}
	if got, want := fizzbuzzTopMsg.CountRequest, int64(2); got != want {
		t.Fatalf("Wrong version return, got %d but want %d", got, want)
	}
	if got, want := fizzbuzzTopMsg.FizzBuzzParams.Str1, "TWO_CALL"; got != want {
		t.Fatalf("Wrong version return, got %s but want %s", got, want)
	}
}
