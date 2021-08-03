// Package middlewares provides some middlewares and a convenient way to chain it.
package middlewares

import "net/http"

// Middleware type
type Middleware func(http.Handler) http.Handler

// UseMiddleware is an helper to declare the middlewares chain.
func useMiddlewareHandler(h http.Handler, mw ...Middleware) http.Handler {
	for i := range mw {
		h = mw[len(mw)-1-i](h)
	}

	return h
}

// UseMiddleware provides the ability to declare the middlewares chain.
func UseMiddleware(mw ...Middleware) func(http.HandlerFunc) http.Handler {
	return func(fn http.HandlerFunc) http.Handler {
		return useMiddlewareHandler(fn, mw...)
	}
}
