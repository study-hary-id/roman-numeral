package utils

import (
	"encoding/json"
	"net/http"

	"github.com/study-hary-id/roman-numeral-api/models"
)

// ConvertToRoman return roman numeral effectively until 3999.
func ConvertToRoman(num int) (romanNum string) {

	var (
		keyNumerals = [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
		index       = len(keyNumerals) - 1
	)
	for num != 0 {
		maxChance := num / keyNumerals[index]
		num %= keyNumerals[index]

		for maxChance != 0 {
			romanNum += Numerals[keyNumerals[index]]
			maxChance -= 1
		}
		index -= 1
	}
	return romanNum

}

// ResponseJSON makes the response with models.Payload as JSON format.
func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, jsonErr := json.Marshal(payload)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ResponseError makes the error response with models.ErrorPayload as json format.
func ResponseError(w http.ResponseWriter, status int, payload models.Errors) {
	ResponseJSON(w, status, models.ErrorPayload{Errors: payload})
}
