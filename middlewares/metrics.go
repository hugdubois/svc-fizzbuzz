package middlewares

import (
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fizzbuzz_http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
		},
		[]string{"code", "method", "path"},
	)
	latencyHistogram = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "fizzbuzz_http_requests_duration_millisecond",
			Help: "How long it took to process the request, partitioned by status code, method and HTTP path.",
		},
		[]string{"code", "method", "path"},
	)
	inFlightGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "fizzbuzz_http_requests_active",
			Help: "How many active in-flight requests, partitioned by status code, method, HTTP path.",
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(latencyHistogram)
	prometheus.MustRegister(inFlightGauge)
}

// NewMetrics returns a new metrics middleware (here this is prometheus).
func NewMetrics() Middleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			inFlightGauge.WithLabelValues(r.Method, r.URL.Path).Inc()
			metrics := httpsnoop.CaptureMetrics(next, w, r)

			requestCounter.WithLabelValues(http.StatusText(metrics.Code), r.Method, r.URL.Path).Inc()

			inFlightGauge.WithLabelValues(r.Method, r.URL.Path).Dec()
			latencyHistogram.WithLabelValues(http.StatusText(metrics.Code), r.Method, r.URL.Path).Observe(float64(metrics.Duration / time.Millisecond))
		}

		return http.HandlerFunc(fn)
	}
}
