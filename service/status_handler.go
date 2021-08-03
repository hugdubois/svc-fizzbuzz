package service

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// StatusHandler is a http handler that returns the service status.
//
// @Summary service status
//
// @Description returns the service status
//
// @Produce json
// @Success 200 {object} StatusResponse
// @Failure 500 {object} ErrorMessage
//
// @Router /status [get]
func (svc Service) StatusHandler(w http.ResponseWriter, r *http.Request) {
	pong, err := pingStore()
	if err != nil {
		log.Errorf("store ping failure - %s", err.Error())
	}

	msg := StatusResponse{
		SvcAlive:   true,
		StoreAlive: pong == "PONG",
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
