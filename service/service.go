// service package is the simple fizzbuzz service
package service

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/hugdubois/svc-fizzbuzz/middlewares"
	"github.com/hugdubois/svc-fizzbuzz/store"
)

var (
	name    = "svc-fizzbuzz"
	version = "latest" // injected with -ldflags in Makefile

	// this is useful for the testing
	pingStore = store.Ping

	// fizzbuzz hits stat
	fizzbuzzHits store.Hitable
)

func init() {
	fizzbuzzHits = store.NewHits("fizzbuzz")
}

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
	router.Handle("/api/v1/fizzbuzz", useMiddlewares(svc.FizzBuzzHandler))
	router.Handle("/api/v1/fizzbuzz/top", useMiddlewares(svc.FizzBuzzTopHandler))

	// prometheus metrics handler
	router.Handle("/metrics", promhttp.Handler())

	// welcome msg on / else return a 404
	router.Handle("/", useMiddlewares(svc.IndexHandler))

	return router
}
