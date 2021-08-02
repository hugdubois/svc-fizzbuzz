// service package is the simple fizzbuzz service
package service

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	"github.com/hugdubois/svc-fizzbuzz/service/handlers"
	"github.com/hugdubois/svc-fizzbuzz/service/middlewares"
	"github.com/hugdubois/svc-fizzbuzz/store"
)

var (
	name    = "svc-fizzbuzz"
	version = "latest" // injected with -ldflags in Makefile

	// this is useful for the testing
	pingStore = store.Ping
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
func (svc Service) NewRouter(corsOrigin string) *http.ServeMux {
	router := http.NewServeMux()

	useMiddlewares := middlewares.UseMiddleware(
		// panic recover
		middlewares.NewRecover(),
		// prometheus instrument handler it's must be at the top of middleware chain
		middlewares.NewMetrics(),
		// nice log with metrics
		middlewares.NewLogging(name),
		// allow cors origin
		middlewares.NewCors(corsOrigin),
	)

	// basic endpoints
	router.Handle("/status", useMiddlewares(svc.StatusHandler))
	router.Handle("/version", useMiddlewares(svc.VersionHandler))

	// service api endpoints
	router.Handle("/api/v1/fizzbuzz", useMiddlewares(handlers.FizzBuzzHandler))
	router.Handle("/api/v1/fizzbuzz/top", useMiddlewares(handlers.FizzBuzzTopHandler))

	// prometheus metrics handler
	router.Handle("/metrics", promhttp.Handler())

	// welcome msg on / else return a 404
	router.Handle("/", useMiddlewares(func(w http.ResponseWriter, r *http.Request) {
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
	SvcAlive   bool `json:"svc-alive"`
	StoreAlive bool `json:"store-alive"`
}

// StatusHandler is a http handler which returns the service status
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
