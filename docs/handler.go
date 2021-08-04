// Package docs provides the swagger definition of the service.
package docs

import (
	"fmt"
	"net/http"
)

// Handler is a http handler that returns swagger definition of the service
func Handler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		swag := &s{}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, swag.ReadDoc())
	}

	return http.HandlerFunc(fn)
}
