// service package is the simple fizzbuzz service
package service

import (
	"encoding/json"
	"net/http"

	"github.com/hugdubois/svc-fizzbuzz/service/handlers"
	"github.com/hugdubois/svc-fizzbuzz/service/middlewares"
)

var (
	name    = "svc-fizzbuzz"
	version = "latest" // injected with -ldflags in Makefile
)

// Service is the fizzbuzz service
type Service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// NewService return new fizzbuzz service
func NewService() *Service {
	return &Service{
		Name:    name,
		Version: version,
	}
}

// NewRouter provides a http.serverMux
func (svc Service) NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	useMiddleware := middlewares.UseMiddleware(
		// panic recover
		middlewares.RecoverMiddleware,
		// nice log with metrics
		middlewares.NewLoggingMiddleware(name),
	)

	router.Handle("/status", useMiddleware(svc.StatusHandler))
	router.Handle("/version", useMiddleware(svc.VersionHandler))

	// welcome msg on / else return a 404
	router.Handle("/", useMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// The "/" pattern matches everything not matched by previous handlers
		if r.URL.Path != "/" {
			handlers.ErrorHandler(w, r, http.StatusNotFound, "Not Found")
			return
		}
		svc.VersionHandler(w, r)
	}))

	return router
}

// BASIC SERVICE HANDLERS

// StatusResponse is the message returned by Status handler
type StatusResponse struct {
	Alive bool `json:"alive"`
}

// StatusHandler is a http handler which returns the service status
func (svc Service) StatusHandler(w http.ResponseWriter, r *http.Request) {
	msg := StatusResponse{
		Alive: true,
	}

	output, err := json.Marshal(msg)
	if err != nil {
		handlers.ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// VersionHandler is a http handler which returns the service version
func (svc Service) VersionHandler(w http.ResponseWriter, r *http.Request) {
	output, err := json.Marshal(svc)
	if err != nil {
		handlers.ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
