package handlers

import (
	"net/http"

	"github.com/study-hary-id/roman-numeral-api/models"
)

// ZeroParamHandler handle roman-numbers endpoint with number 0 parameter.
func ZeroParamHandler(w http.ResponseWriter) {
	errJson := models.Errors{
		Status: http.StatusNotFound,
		Title:  "Cannot Convert 0",
		Detail: "Roman numerals do not have the number 0.",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	_, err := w.Write(marshalJsonHandler(
		w, models.ErrorPayload{Errors: errJson},
	))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
