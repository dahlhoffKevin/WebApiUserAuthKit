package errorhandler

import (
	"encoding/json"
	"errors"
	"net/http"
)

type HTTPError struct {
	Code    int    `json:"-"` // code gehört in den header und nicht ins json
	Message string `json:"error"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func New(code int, message string) *HTTPError {
	if message == "" {
		message = http.StatusText(code)
	}
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

// Common helpers
func Unauthorized() *HTTPError {
	return New(http.StatusUnauthorized, "unauthorized")
}

func Forbidden() *HTTPError {
	return New(http.StatusForbidden, "forbidden")
}

func NotFound() *HTTPError {
	return New(http.StatusNotFound, "not found")
}

func BadRequest(message string) *HTTPError {
	return New(http.StatusBadRequest, message)
}

func Internal() *HTTPError {
	return New(http.StatusInternalServerError, "internal server error")
}

func Write(w http.ResponseWriter, err error) {
	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		httpErr = Internal()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpErr.Code)
	_ = json.NewEncoder(w).Encode(httpErr)
}
