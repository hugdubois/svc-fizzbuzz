// Package handlers provides the http handlers.
package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hugdubois/svc-fizzbuzz/core"
	"github.com/hugdubois/svc-fizzbuzz/store"
)

var fizzbuzzHits store.Hitable

func init() {
	fizzbuzzHits = store.NewHits("fizzbuzz")
}

// parseIntValue parse an expected int parameter
// if the parameter is missing a default value is returned
// if the parameter is not an int an error is returned
func parseIntValue(r *http.Request, name string, dVal int) (int, error) {
	val := r.FormValue(name)
	if val == "" {
		return dVal, nil
	}

	iVal, err := strconv.Atoi(val)
	if err != nil {
		return dVal, fmt.Errorf("Bad parameter: '%s' must be a positive number - got (%s)", name, val)
	}
	return iVal, nil
}

// parseStringValue parse an expected string parameter
// if the parameter is missing a default value is returned
func parseStringValue(r *http.Request, name string, dVal string) string {
	val := r.FormValue(name)
	if val == "" {
		return dVal
	}

	return val
}

// parseFizzbuzzParams parse request parameters and returns a FizzBuzzParams
// if an error occured an error is returned
func parseFizzbuzzParams(r *http.Request) (*core.FizzBuzzParams, error) {
	p := core.DefaultFizzBuzzParams

	l, err := parseIntValue(r, "limit", p.Limit)
	if err != nil {
		return nil, err
	}
	p.Limit = l

	m1, err := parseIntValue(r, "int1", p.Int1)
	if err != nil {
		return nil, err
	}
	p.Int1 = m1

	m2, err := parseIntValue(r, "int2", p.Int2)
	if err != nil {
		return nil, err
	}
	p.Int2 = m2

	p.Str1 = parseStringValue(r, "str1", p.Str1)
	p.Str2 = parseStringValue(r, "str2", p.Str2)

	return &p, nil
}

// encodeFizzbuzzParams encodes a FizzBuzzParams to a query string
func encodeFizzbuzzParams(params core.FizzBuzzParams) string {
	qsParams := url.Values{}
	qsParams.Add("limit", strconv.Itoa(params.Limit))
	qsParams.Add("int1", strconv.Itoa(params.Int1))
	qsParams.Add("str1", params.Str1)
	qsParams.Add("int2", strconv.Itoa(params.Int2))
	qsParams.Add("str2", params.Str2)

	return qsParams.Encode()
}

// fizzbuzzHit add a fizzbuzz hit from a FizzBuzzParams
func fizzbuzzHit(params core.FizzBuzzParams) {
	fizzbuzzHits.Add(encodeFizzbuzzParams(params), 1)
}

// fizzbuzzTopHit retreives the top fizzbuzz hits
func fizzbuzzTopHit() (*core.FizzBuzzParams, int64, error) {
	top, count, err := fizzbuzzHits.Top()
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("GET", "?"+top, nil)
	if err != nil {
		return nil, 0, err
	}

	params, err := parseFizzbuzzParams(req)
	if err != nil {
		return nil, 0, err
	}

	return params, count, nil
}
