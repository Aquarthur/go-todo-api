package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Aquarthur/go-todo-api/domain"
)

const errJSONFormat = `{
	"error": "%s"
}`

func handleError(err error, w http.ResponseWriter) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf(errJSONFormat, err.Error())))
	case errors.Is(err, domain.ErrConflict):
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(fmt.Sprintf(errJSONFormat, err.Error())))
	default:
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
