package helpers

import (
	"encoding/json"
	"net/http"
)

// SendSuccessResponse sends a JSON response for successful operations
func SendSuccessResponse(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := SuccessResponse{
		Data:       data,
		Message:    message,
		StatusCode: statusCode,
	}
	json.NewEncoder(w).Encode(response)
}

// SendErrorResponse sends a JSON response for errors
func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
	}
	json.NewEncoder(w).Encode(response)
}
