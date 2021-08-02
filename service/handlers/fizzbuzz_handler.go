package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hugdubois/svc-fizzbuzz/core"
)

// FizzBuzzResponse is the message returned by FizzBuzz handler
type FizzBuzzResponse struct {
	FizzBuzz []string `json:"fizzbuzz"`
}

// FizzBuzzHandler is a http handler which returns the FizzBuzz core function
func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parseFizzbuzzParams(r)

	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fizzbuzz, err := core.FizzBuzz(*params)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	msg := FizzBuzzResponse{
		FizzBuzz: fizzbuzz,
	}

	output, err := json.Marshal(msg)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// hit fizzbuzz request
	fizzbuzzHit(*params)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}