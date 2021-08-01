package handlers

import (
	"encoding/json"
	"net/http"
)

// ErrorMessage is the message returned when an error has occurred
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorHandler is the http handler wich returns an error
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, msg string) {
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
