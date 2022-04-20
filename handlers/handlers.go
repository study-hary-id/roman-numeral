package handlers

import (
	"fmt"
	"net/http"

	"github.com/study-hary-id/roman-numeral-api/models"
	"github.com/study-hary-id/roman-numeral-api/utils"
)

// NotFound handles wrong URLs/routes/endpoints.
func NotFound(w http.ResponseWriter, r *http.Request) {
	errJSON := models.Errors{
		Status: http.StatusNotFound,
		Title:  "Not Found",
		Detail: fmt.Sprintf(`%q not found.`, r.URL.Path),
	}
	utils.ResponseError(w, http.StatusNotFound, errJSON)
}

// RomanNumber handles roman-numbers endpoint with correct parameter.
func RomanNumber(w http.ResponseWriter, _ *http.Request, num int) {
	resJSON := models.Numeral{
		Roman:   utils.ConvertToRoman(num),
		Ordinal: num,
	}
	utils.ResponseSuccess(w, http.StatusOK, resJSON)
}
