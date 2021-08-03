package service

import (
	"encoding/json"
	"net/http"
)

// FizzBuzzTopHandler is a http handler that returns the most used request of
// the fizzbuzz endpoint call with it's parameters.
//
// @Summary Most used /api/v1/fizzbuzz request usage statistics.
//
// @description Returns usage statistics of the /api/v1/fizzbuzz endpoint.
// @description It allows the users to know what the number of hits of that endpoint.
// @description And returns the parameters corresponding to it.
//
// @Produce json
// @Success 200 {object} FizzBuzzTopResponse
// @Failure 500 {object} ErrorMessage
//
// @Router /api/v1/fizzbuzz/top [get]
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
