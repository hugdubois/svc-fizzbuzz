package service

import (
	"encoding/json"
	"net/http"
)

// VersionHandler is a http handler that returns the service version.
//
// @Summary Service version
//
// @Description Returns the service version
//
// @Produce json
// @Success 200 {object} Service
// @Failure 500 {object} ErrorMessage
//
// @Router /version [get]
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
