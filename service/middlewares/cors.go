package middlewares

import "net/http"

// NewCorsMiddleware return a middleware that allow request to be made through origin
func NewCorsMiddleware(origin string) Middleware {
	if origin != "" {
		return func(h http.Handler) http.Handler {
			fn := func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS")
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

				if r.Method == "OPTIONS" {
					return
				}

				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				if r.Method == "HEAD" {
					return
				}
				h.ServeHTTP(w, r)
			}

			return http.HandlerFunc(fn)
		}
	}

	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
