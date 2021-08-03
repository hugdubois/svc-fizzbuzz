package service

import (
	"encoding/json"
	"net/http"

	"github.com/hugdubois/svc-fizzbuzz/core"
)

// FizzBuzzHandler is a http handler that returns the result of FizzBuzz core
// function.
//
// @Summary fizzbuzz computation
//
// @description Returns a list of strings with numbers from 1 to `limit`, where:
// @description all multiples of `int1` are replaced by `str1`,
// @description all multiples of `int2` are replaced by `str2`,
// @description all multiples of `int1` and `int2` are replaced by `str1str2`.
//
// @Param limit query int false "fizzbuzz from 1 to limit"
// @Param int1 query int false "multiples replaced 1"
// @Param str1 query string false "replacement string 1"
// @Param int2 query int false "multiples replaced 2"
// @Param str2 query string false "replacement string 2"
//
// @Produce json
// @Success 200 {object} FizzBuzzResponse
// @Failure 422,500 {object} ErrorMessage
//
// @Router /api/v1/fizzbuzz [get]
func (svc Service) FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parseFizzbuzzParams(r)

	if err != nil {
		svc.ErrorHandler(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	fizzbuzz, err := core.FizzBuzz(*params)
	if err != nil {
		svc.ErrorHandler(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	msg := FizzBuzzResponse{
		FizzBuzz: fizzbuzz,
	}

	output, err := json.Marshal(msg)
	if err != nil {
		svc.ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// hit fizzbuzz request
	fizzbuzzHit(*params)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
