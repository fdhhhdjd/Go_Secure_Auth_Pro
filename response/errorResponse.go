package response

import (
	"time"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Now     int64  `json:"now"`
}

// NewErrorResponse creates a new ErrorResponse
func NewErrorResponse(message string, status int) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Status:  status,
		Now:     time.Now().Unix(),
	}
}

// BadRequestError represents a 400 Bad Request error
func BadRequestError(messages ...string) *ErrorResponse {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusBadRequest)
	}
	return NewErrorResponse(message, StatusBadRequest)
}

// NotFoundError represents a 404 Not Found error
func NotFoundError(messages ...string) *ErrorResponse {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusNotFound)
	}
	return NewErrorResponse(message, StatusNotFound)
}

// UnauthorizedError represents a 401 Unauthorized error
func UnauthorizedError(messages ...string) *ErrorResponse {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusUnauthorized)
	}
	return NewErrorResponse(message, StatusUnauthorized)
}

func ForbiddenError(messages ...string) *ErrorResponse {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusForbidden)
	}
	return NewErrorResponse(message, StatusForbidden)
}

// InternalServerError represents a 500 Internal Server Error
func InternalServerError(messages ...string) *ErrorResponse {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusInternalServerError)
	}
	return NewErrorResponse(message, StatusInternalServerError)
}

// ServiceUnavailable represents a 503 Service Unavailable
func ServiceUnavailable(messages ...string) *ErrorResponse {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusServiceUnavailable)
	}
	return NewErrorResponse(message, StatusServiceUnavailable)
}
