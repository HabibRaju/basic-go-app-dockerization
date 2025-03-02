package dto

import (
	"encoding/json"
	"net/http"
)

// Response is the common response structure
type Response struct {
	Status     string      `json:"status"`      // "success" or "failed"
	StatusCode int         `json:"status_code"` // HTTP status code (e.g., 200, 400)
	Message    string      `json:"message"`     // Descriptive message
	Data       interface{} `json:"data"`        // Main response data
	Metadata   interface{} `json:"_metadata"`   // Additional metadata (e.g., pagination)
}

// SuccessResponse creates a success response
func SuccessResponse(statusCode int, message string, data interface{}, metadata interface{}) *Response {
	return &Response{
		Status:     "success",
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Metadata:   metadata,
	}
}

// ErrorResponse creates an error response
func ErrorResponse(statusCode int, message string) *Response {
	return &Response{
		Status:     "failed",
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
		Metadata:   nil,
	}
}

// WriteJSON writes the response as JSON with optional headers
func WriteJSON(w http.ResponseWriter, statusCode int, response *Response, headers map[string]string) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")
	for key, value := range headers {
		w.Header().Set(key, value)
	}

	// Write status code and JSON response
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
