package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rgamz/go-products-api/internal/models"
)

// WriteJSON sets the "Content-Type" header to "application/json",
// transmits the specified HTTP status code,
// and serializes value v into the response body.
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// WriteError transmits a standardized JSON erro response using models.APIError,
// mapping the provided HTTP status code and error message.
func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, models.APIError{
		Code:    status,
		Message: message,
	})
}
