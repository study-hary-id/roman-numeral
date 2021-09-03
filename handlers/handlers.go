package handlers

import (
	"fmt"
	"net/http"

	"github.com/study-hary-id/roman-numeral-api/models"
	"github.com/study-hary-id/roman-numeral-api/utils"
)

// WrongRouteHandler handle wrong URLs/routes/endpoints.
func WrongRouteHandler(w http.ResponseWriter, r *http.Request) {
	errJSON := models.Errors{
		Status: http.StatusBadRequest,
		Title:  "Wrong URL",
		Detail: fmt.Sprintf("%s URL not found.", r.URL.Path),
	}
	utils.ResponseError(w, http.StatusBadRequest, errJSON)
}

// RomanNumberHandler handle roman-numbers endpoint with correct parameter.
func RomanNumberHandler(w http.ResponseWriter, num int) {
	resJSON := models.Numeral{
		Roman:   utils.ConvertToRoman(num),
		Ordinal: num,
	}
	utils.ResponseJSON(w, http.StatusOK, models.Payload{Data: resJSON})
}

// IncompleteRouteHandler handle right URLs/routes/endpoints but missing parameter.
func IncompleteRouteHandler(w http.ResponseWriter) {
	errJSON := models.Errors{
		Status: http.StatusBadRequest,
		Title:  "Incomplete URL",
		Detail: "/roman-numbers requires numeric parameter.",
	}
	utils.ResponseError(w, http.StatusBadRequest, errJSON)
}

// WrongParamHandler handler right URLs/routes/endpoints but wrong parameter.
func WrongParamHandler(w http.ResponseWriter) {
	errJSON := models.Errors{
		Status: http.StatusBadRequest,
		Title:  "Wrong Parameter",
		Detail: "/roman-numbers requires numeric parameter.",
	}
	utils.ResponseError(w, http.StatusBadRequest, errJSON)
}

// ZeroParamHandler handle roman-numbers endpoint with number 0 parameter.
func ZeroParamHandler(w http.ResponseWriter) {
	errJSON := models.Errors{
		Status: http.StatusNotFound,
		Title:  "Number 0 Not Found",
		Detail: "Roman numerals do not have the number 0.",
	}
	utils.ResponseError(w, http.StatusNotFound, errJSON)
}
