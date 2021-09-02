package models

// Numeral contains roman numeral and ordinary numeral.
type Numeral struct {
	Roman   string `json:"roman"`
	Ordinal int    `json:"ordinal"`
}

// Errors contains a set of data for error JSON response.
type Errors struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// Payload contains response data for http request.
type Payload struct {
	Data Numeral `json:"data"`
}

// ErrorPayload contains response error for http request.
type ErrorPayload struct {
	Errors `json:"errors"`
}
