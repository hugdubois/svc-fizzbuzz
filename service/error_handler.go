package service

import (
	"encoding/json"
	"net/http"
)

// ErrorHandler is the http handler that returns an error.
func (svc Service) ErrorHandler(w http.ResponseWriter, r *http.Request, status int, msg string) {
	eMsg := ErrorMessage{
		Code:    status,
		Message: msg,
	}
	output, err := json.Marshal(eMsg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
