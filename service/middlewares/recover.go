package middlewares

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// RecoverMiddleware is a middleware to recover the errors
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defer the recovery
		defer func() {
			if r := recover(); r != nil {
				log.Printf("There was a panic with this message: %v\n", r)
				switch x := r.(type) {
				case string:
					http.Error(w, r.(string), http.StatusInternalServerError)
				case error:
					err := x
					http.Error(w, err.Error(), http.StatusInternalServerError)
				default:
					http.Error(w, "unknown panic", http.StatusInternalServerError)
				}
			}
		}()

		// call the actual api handler
		next.ServeHTTP(w, r)
	})
}
