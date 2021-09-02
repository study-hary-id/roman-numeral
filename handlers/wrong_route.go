package handlers

import (
	"fmt"
	"net/http"

	"github.com/study-hary-id/roman-numeral-api/models"
)

// WrongRouteHandler handle wrong URLs/routes/endpoints.
func WrongRouteHandler(w http.ResponseWriter, r *http.Request) {
	errJson := models.Errors{
		Status: http.StatusBadRequest,
		Title:  "Wrong URL",
		Detail: fmt.Sprintf("%s URL not found.", r.URL.Path),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	_, err := w.Write(marshalJsonHandler(
		w, models.ErrorPayload{Errors: errJson},
	))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
