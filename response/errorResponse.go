package response

import (
	"time"

	"github.com/gin-gonic/gin"
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

// Send sends the error response to the client.
// It aborts the request and responds with the error response as JSON.
func (sr *ErrorResponse) Send(c *gin.Context) {
	c.AbortWithStatusJSON(sr.Status, sr)
}

// BadRequestError represents a 400 Bad Request error
func BadRequestError(c *gin.Context, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusBadRequest)
	}
	response := NewErrorResponse(message, StatusBadRequest)
	response.Send(c)

}

// NotFoundError represents a 404 Not Found error
func NotFoundError(c *gin.Context, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusNotFound)
	}
	response := NewErrorResponse(message, StatusNotFound)
	response.Send(c)

}

// TooManyRequestsError represents a 429 Too Many Requests error
func TooManyRequestsError(c *gin.Context, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusTooManyRequests)
	}
	response := NewErrorResponse(message, StatusTooManyRequests)
	response.Send(c)
}

// UnauthorizedError represents a 401 Unauthorized error
func UnauthorizedError(c *gin.Context, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusUnauthorized)
	}
	response := NewErrorResponse(message, StatusUnauthorized)
	response.Send(c)
}

// ForbiddenError handles the generation and sending of a Forbidden error response.
func ForbiddenError(c *gin.Context, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusForbidden)
	}
	response := NewErrorResponse(message, StatusForbidden)
	response.Send(c)
}

// EntityTooLargeError handles the HTTP 413 Request Entity Too Large error.
func EntityTooLargeError(c *gin.Context, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusRequestEntityTooLarge)
	}
	response := NewErrorResponse(message, StatusRequestEntityTooLarge)
	response.Send(c)
}

// UnSupportMediaTypeError handles the unsupported media type error by sending an error response to the client.
func UnSupportMediaTypeError(c *gin.Context, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusUnsupportedMediaType)
	}
	response := NewErrorResponse(message, StatusUnsupportedMediaType)
	response.Send(c)
}

// InternalServerError represents a 500 Internal Server Error
func InternalServerError(c *gin.Context, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusInternalServerError)
	}
	response := NewErrorResponse(message, StatusInternalServerError)
	response.Send(c)

}

// ServiceUnavailable represents a 503 Service Unavailable
func ServiceUnavailable(c *gin.Context, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusServiceUnavailable)
	}
	response := NewErrorResponse(message, StatusServiceUnavailable)
	response.Send(c)
}
