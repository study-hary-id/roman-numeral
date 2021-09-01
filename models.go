package main

// numeral contains roman numeral and ordinary numeral.
type numeral struct {
	Roman   string `json:"roman"`
	Ordinal int    `json:"ordinal"`
}

// errors contains a set of data for error JSON response.
type errors struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// payload contains response data for http request.
type payload struct {
	Data numeral `json:"data"`
}

// errorPayload contains response error for http request.
type errorPayload struct {
	errors `json:"errors"`
}
