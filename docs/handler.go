package docs

import (
	"fmt"
	"net/http"
)

// Handler is a http handler that returns swagger definition of the service
func Handler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		sw := &s{}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, sw.ReadDoc())
	}

	return http.HandlerFunc(fn)
}
