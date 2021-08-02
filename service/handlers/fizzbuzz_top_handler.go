package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hugdubois/svc-fizzbuzz/core"
)

// FizzBuzzTopResponse is the message returned by FizzBuzzTopHandler
type FizzBuzzTopResponse struct {
	Data FizzBuzzTopResponseData `json:"data"`
}

// FizzBuzzTopResponseData is the core of message returned by FizzBuzzTopHandler
type FizzBuzzTopResponseData struct {
	Params       core.FizzBuzzParams `json:"params"`
	CountRequest int64               `json:"count_request"`
}

// FizzBuzzTopHandler is a http handler which returns the most used request of
// the fizzbuzz endpoint call and it's parameters
func FizzBuzzTopHandler(w http.ResponseWriter, r *http.Request) {
	params, countReq, err := fizzbuzzTopHit()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	msg := FizzBuzzTopResponse{
		Data: FizzBuzzTopResponseData{
			Params:       *params,
			CountRequest: countReq,
		},
	}

	output, err := json.Marshal(msg)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
