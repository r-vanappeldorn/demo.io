package httperrors

type Error struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type Errors []Error

type ErrorRes struct {
	Errors     Errors `json:"errors"`
	StatusCode int    `json:"statusCode"`
}
