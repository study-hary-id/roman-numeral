package handlers

import (
	"net/http"

	"github.com/study-hary-id/roman-numeral-api/models"
	"github.com/study-hary-id/roman-numeral-api/utils"
)

// RomanNumberHandler handle roman-numbers endpoint with correct parameter.
func RomanNumberHandler(w http.ResponseWriter, num int) {
	resJson := models.Numeral{
		Roman:   utils.ConvertToRoman(num),
		Ordinal: num,
	}
	w.Header().Set("Content-Type", "application/json")

	_, err := w.Write(marshalJsonHandler(
		w, models.Payload{Data: resJson},
	))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
