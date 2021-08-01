package core

import (
	"fmt"
	"reflect"
	"testing"
)

// assertInvalidFizzBuzz is a helper function to test the invalid parameters
func assertInvalidFizzBuzz(t *testing.T, params FizzBuzzParams, errMsg string) {
	t.Helper()
	_, err := FizzBuzz(params)
	if err == nil {
		t.Fatalf("Got no error when \"%s\" error message is expected - (%#v)", errMsg, params)
	}
	if err.Error() != errMsg {
		t.Fatalf("Got unexpected error \"%s\" when \"%s\" error message is expected - (%#v)", err.Error(), errMsg, params)
	}
}

// assertValidFizzBuzz is a helper function to test the valid parameters
func assertValidFizzBuzz(t *testing.T, params FizzBuzzParams, expected []string) {
	t.Helper()
	res, err := FizzBuzz(params)
	if err != nil {
		t.Fatalf("Got an unexpected error: %s - (%#v)", err.Error(), params)
	}
	if got := reflect.DeepEqual(expected, res); !got {
		t.Fatalf("%s is not equal to expected (%s)", res, expected)
	}
	if len(res) != params.Limit {
		t.Fatalf("res length %d is not equal to expected length (%d)", len(res), params.Limit)
	}
}

// TestValidFizzBuzzParams is a function to test the FizzBuzz function with some valid parameters
func Test_ValidFizzBuzzParams(t *testing.T) {
	//default params
	assertValidFizzBuzz(
		t,
		DefaultFizzBuzzParams,
		[]string{
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
	)
	assertValidFizzBuzz(
		t,
		FizzBuzzParams{Limit: 70, Int1: 7, Str1: "bon", Int2: 9, Str2: "coin"},
		[]string{
			"1", "2", "3", "4", "5", "6", "bon", "8", "coin", "10",
			"11", "12", "13", "bon", "15", "16", "17", "coin", "19", "20",
			"bon", "22", "23", "24", "25", "26", "coin", "bon", "29", "30",
			"31", "32", "33", "34", "bon", "coin", "37", "38", "39", "40",
			"41", "bon", "43", "44", "coin", "46", "47", "48", "bon", "50",
			"51", "52", "53", "coin", "55", "bon", "57", "58", "59", "60",
			"61", "62", "boncoin", "64", "65", "66", "67", "68", "69", "bon",
		},
	)
}

// TestInvalidFizzBuzzParams is a function to test the FizzBuzz function with some invalid parameters
func Test_InvalidFizzBuzzParams(t *testing.T) {
	assertInvalidFizzBuzz(
		t,
		FizzBuzzParams{Limit: 0, Int1: 7, Str1: "bon", Int2: 9, Str2: "coin"},
		"Bad parameter: 'limit' must be a positive number - got (0)",
	)
	assertInvalidFizzBuzz(
		t,
		FizzBuzzParams{Limit: -1, Int1: 7, Str1: "bon", Int2: 9, Str2: "coin"},
		"Bad parameter: 'limit' must be a positive number - got (-1)",
	)
	assertInvalidFizzBuzz(
		t,
		FizzBuzzParams{Limit: 10, Int1: 0, Str1: "bon", Int2: 9, Str2: "coin"},
		"Bad parameter: multiple must be a positive number - got (int1: 0)",
	)
	assertInvalidFizzBuzz(
		t,
		FizzBuzzParams{Limit: 10, Int1: -1, Str1: "bon", Int2: 9, Str2: "coin"},
		"Bad parameter: multiple must be a positive number - got (int1: -1)",
	)
	assertInvalidFizzBuzz(
		t,
		FizzBuzzParams{Limit: 10, Int1: 7, Str1: "bon", Int2: 0, Str2: "coin"},
		"Bad parameter: multiple must be a positive number - got (int2: 0)",
	)
	assertInvalidFizzBuzz(
		t,
		FizzBuzzParams{Limit: 10, Int1: 7, Str1: "bon", Int2: -1, Str2: "coin"},
		"Bad parameter: multiple must be a positive number - got (int2: -1)",
	)
	assertInvalidFizzBuzz(
		t,
		FizzBuzzParams{Limit: FizzBuzzLimitMax + 1, Int1: 7, Str1: "bon", Int2: 9, Str2: "coin"},
		fmt.Sprintf(
			"Bad parameter: maximum 'limit' reached - got (%d) max (%d)",
			FizzBuzzLimitMax+1,
			FizzBuzzLimitMax,
		),
	)
}
