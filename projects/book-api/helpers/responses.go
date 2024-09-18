package helpers

type SuccessResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
	StatusCode int         `json:"statusCode"`
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
