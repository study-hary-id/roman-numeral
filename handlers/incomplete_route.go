package handlers

import (
	"net/http"

	"github.com/study-hary-id/roman-numeral-api/models"
)

// IncompleteRouteHandler handle right URLs/routes/endpoints but missing parameter.
func IncompleteRouteHandler(w http.ResponseWriter) {
	errJson := models.Errors{
		Status: http.StatusBadRequest,
		Title:  "Incomplete Endpoint",
		Detail: "/roman-numbers requires numeric parameter.",
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
