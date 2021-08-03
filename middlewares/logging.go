package middlewares

import (
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	log "github.com/sirupsen/logrus"
)

// loggingMiddlewareBefore performs log before the request execution.
func loggingMiddlewareBefore(entry *log.Entry, r *http.Request, t time.Time) *log.Entry {
	// Try to get the request ID (correlation ID)
	if reqID := r.Header.Get("X-Request-Id"); reqID != "" {
		entry = entry.WithField("request_id", reqID)
	}

	// Try to get the forwarded for header
	if forward := r.Header.Get("X-Forwarded-For"); forward != "" {
		entry = entry.WithField("forwarded_for", forward)
	}

	return entry.WithFields(log.Fields{
		"uri":         r.RequestURI,
		"proto":       r.Proto,
		"method":      r.Method,
		"host":        r.Host,
		"user_agent":  r.UserAgent(),
		"reqest_size": r.ContentLength,
		"time":        t.Format(time.RFC3339),
	})
}

// loggingMiddlewareAfter performs log after the request execution.
func loggingMiddlewareAfter(entry *log.Entry, metrics httpsnoop.Metrics, name string) *log.Entry {
	return entry.WithFields(log.Fields{
		"status":        metrics.Code,
		"text_status":   http.StatusText(metrics.Code),
		"took":          metrics.Duration,
		"response_size": metrics.Written,
	})
}

// NewLogging returns a new logging middleware.
func NewLogging(name string) Middleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			entry := log.NewEntry(log.StandardLogger())

			entry = loggingMiddlewareBefore(entry, r, start)
			entry.Info("started handling request")
			metrics := httpsnoop.CaptureMetrics(next, w, r)
			loggingMiddlewareAfter(entry, metrics, name).Info("completed handling request")
		}

		return http.HandlerFunc(fn)
	}
}
