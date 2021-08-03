package service

import "net/http"

// IndexHandler is the default handler when the url matches anything
func (svc Service) IndexHandler(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything not matched by previous handlers
	if r.URL.Path != "/" {
		svc.ErrorHandler(w, r, http.StatusNotFound, "Not Found")
		return
	}
	svc.VersionHandler(w, r)
}
