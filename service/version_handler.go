package service

import (
	"encoding/json"
	"net/http"
)

// VersionHandler is a http handler which returns the service version
func (svc Service) VersionHandler(w http.ResponseWriter, r *http.Request) {
	output, err := json.Marshal(svc)
	if err != nil {
		svc.ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
