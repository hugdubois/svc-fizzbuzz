package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hugdubois/svc-fizzbuzz/helpers"
)

func init() {
	fizzbuzzHits = helpers.NewMockHits("fizzbuzz")
}

// assertValidFizzBuzz is a helper function to test the valid parameters
func assertValidFizzBuzz(t *testing.T, url string, expected FizzBuzzResponse) {
	var fizzbuzzMsg FizzBuzzResponse
	t.Helper()

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+url, nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Fatalf("Invalid status code, got %d but want %d", got, want)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Got an error when reading body: %s", err.Error())
	}

	err = json.Unmarshal(data, &fizzbuzzMsg)
	if err != nil {
		t.Fatalf("Got an error when parsing json: %s", err.Error())
	}

	if got := reflect.DeepEqual(expected, fizzbuzzMsg); !got {
		t.Fatalf("%s is not equal to expected (%s)", fizzbuzzMsg, expected)
	}
}

// assertInvalidFizzBuzz is a helper function to test the valid parameters
func assertInvalidFizzBuzz(t *testing.T, url string, expected ErrorMessage) {
	t.Helper()

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+url, nil)
	resp, _ := http.DefaultClient.Do(req)

	if got, want := resp.StatusCode, http.StatusUnprocessableEntity; got != want {
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

	if got := reflect.DeepEqual(expected, errorMsg); !got {
		t.Fatalf("%v is not equal to expected (%v)", errorMsg, expected)
	}
}

func TestFizzBuzzWithoutParams(t *testing.T) {
	http.HandleFunc("/fizzbuzz", FizzBuzzHandler)

	assertValidFizzBuzz(
		t,
		"/fizzbuzz",
		FizzBuzzResponse{
			FizzBuzz: []string{
				"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz",
				"11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz",
				"fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz",
				"31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz",
				"41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz",
				"fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz",
				"61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz",
				"71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz",
				"fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz",
				"91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"},
		},
	)

	assertValidFizzBuzz(
		t,
		"/fizzbuzz?limit=10",
		FizzBuzzResponse{
			FizzBuzz: []string{
				"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz",
			},
		},
	)

	assertValidFizzBuzz(
		t,
		"/fizzbuzz?str1=bon",
		FizzBuzzResponse{
			FizzBuzz: []string{
				"1", "2", "bon", "4", "buzz", "bon", "7", "8", "bon", "buzz",
				"11", "bon", "13", "14", "bonbuzz", "16", "17", "bon", "19", "buzz",
				"bon", "22", "23", "bon", "buzz", "26", "bon", "28", "29", "bonbuzz",
				"31", "32", "bon", "34", "buzz", "bon", "37", "38", "bon", "buzz",
				"41", "bon", "43", "44", "bonbuzz", "46", "47", "bon", "49", "buzz",
				"bon", "52", "53", "bon", "buzz", "56", "bon", "58", "59", "bonbuzz",
				"61", "62", "bon", "64", "buzz", "bon", "67", "68", "bon", "buzz",
				"71", "bon", "73", "74", "bonbuzz", "76", "77", "bon", "79", "buzz",
				"bon", "82", "83", "bon", "buzz", "86", "bon", "88", "89", "bonbuzz",
				"91", "92", "bon", "94", "buzz", "bon", "97", "98", "bon", "buzz",
			},
		},
	)

	assertValidFizzBuzz(
		t,
		"/fizzbuzz?str2=coin",
		FizzBuzzResponse{
			FizzBuzz: []string{
				"1", "2", "fizz", "4", "coin", "fizz", "7", "8", "fizz", "coin",
				"11", "fizz", "13", "14", "fizzcoin", "16", "17", "fizz", "19", "coin",
				"fizz", "22", "23", "fizz", "coin", "26", "fizz", "28", "29", "fizzcoin",
				"31", "32", "fizz", "34", "coin", "fizz", "37", "38", "fizz", "coin",
				"41", "fizz", "43", "44", "fizzcoin", "46", "47", "fizz", "49", "coin",
				"fizz", "52", "53", "fizz", "coin", "56", "fizz", "58", "59", "fizzcoin",
				"61", "62", "fizz", "64", "coin", "fizz", "67", "68", "fizz", "coin",
				"71", "fizz", "73", "74", "fizzcoin", "76", "77", "fizz", "79", "coin",
				"fizz", "82", "83", "fizz", "coin", "86", "fizz", "88", "89", "fizzcoin",
				"91", "92", "fizz", "94", "coin", "fizz", "97", "98", "fizz", "coin",
			},
		},
	)

	assertValidFizzBuzz(
		t,
		"/fizzbuzz?limit=10&int1=2",
		FizzBuzzResponse{
			FizzBuzz: []string{
				"1", "fizz", "3", "fizz", "buzz", "fizz", "7", "fizz", "9", "fizzbuzz",
			},
		},
	)

	assertValidFizzBuzz(
		t,
		"/fizzbuzz?limit=10&int1=2&int2=3",
		FizzBuzzResponse{
			FizzBuzz: []string{
				"1", "fizz", "buzz", "fizz", "5", "fizzbuzz", "7", "fizz", "buzz", "fizz",
			},
		},
	)

	assertValidFizzBuzz(
		t,
		"/fizzbuzz?limit=70&int1=7&str1=bon&int2=9&str2=coin",
		FizzBuzzResponse{
			FizzBuzz: []string{
				"1", "2", "3", "4", "5", "6", "bon", "8", "coin", "10",
				"11", "12", "13", "bon", "15", "16", "17", "coin", "19", "20",
				"bon", "22", "23", "24", "25", "26", "coin", "bon", "29", "30",
				"31", "32", "33", "34", "bon", "coin", "37", "38", "39", "40",
				"41", "bon", "43", "44", "coin", "46", "47", "48", "bon", "50",
				"51", "52", "53", "coin", "55", "bon", "57", "58", "59", "60",
				"61", "62", "boncoin", "64", "65", "66", "67", "68", "69", "bon",
			},
		},
	)

	assertInvalidFizzBuzz(
		t,
		"/fizzbuzz?limit=INVALID&int1=2&int2=3",
		ErrorMessage{
			Code:    http.StatusUnprocessableEntity,
			Message: "Bad parameter: 'limit' must be a positive number - got (INVALID)",
		},
	)
	assertInvalidFizzBuzz(
		t,
		"/fizzbuzz?int1=INVALID&int2=3",
		ErrorMessage{
			Code:    http.StatusUnprocessableEntity,
			Message: "Bad parameter: 'int1' must be a positive number - got (INVALID)",
		},
	)

	assertInvalidFizzBuzz(
		t,
		"/fizzbuzz?int2=INVALID",
		ErrorMessage{
			Code:    http.StatusUnprocessableEntity,
			Message: "Bad parameter: 'int2' must be a positive number - got (INVALID)",
		},
	)

	assertInvalidFizzBuzz(
		t,
		"/fizzbuzz?int2=-1",
		ErrorMessage{
			Code:    http.StatusUnprocessableEntity,
			Message: "Bad parameter: multiple must be a positive number - got (int2: -1)",
		},
	)
}
