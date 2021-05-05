package handlers

import "net/http"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return new(HealthHandler)
}

func (handler *HealthHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{ \"status\": \"UP\" }"))
}
