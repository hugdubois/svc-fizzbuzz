package service

import (
	"encoding/json"
	"net/http"
)

// FizzBuzzTopHandler is a http handler which returns the most used request of
// the fizzbuzz endpoint call and it's parameters
func (svc Service) FizzBuzzTopHandler(w http.ResponseWriter, r *http.Request) {
	params, countReq, err := fizzbuzzTopHit()
	if err != nil {
		svc.ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
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
		svc.ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
