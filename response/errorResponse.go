package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Now     int64  `json:"now"`
}

// NewErrorResponse creates a new ErrorResponse
func NewErrorResponse(message string, status int, code int) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
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
func BadRequestError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusBadRequest)
	}
	response := NewErrorResponse(message, StatusBadRequest, code)
	response.Send(c)

}

// NotFoundError represents a 404 Not Found error
func NotFoundError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusNotFound)
	}
	response := NewErrorResponse(message, StatusNotFound, code)
	response.Send(c)

}

// TooManyRequestsError represents a 429 Too Many Requests error
func TooManyRequestsError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusTooManyRequests)
	}
	response := NewErrorResponse(message, StatusTooManyRequests, code)
	response.Send(c)
}

// UnauthorizedError represents a 401 Unauthorized error
func UnauthorizedError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusUnauthorized)
	}
	response := NewErrorResponse(message, StatusUnauthorized, code)
	response.Send(c)
}

// ForbiddenError handles the generation and sending of a Forbidden error response.
func ForbiddenError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusForbidden)
	}
	response := NewErrorResponse(message, StatusForbidden, code)
	response.Send(c)
}

// EntityTooLargeError handles the HTTP 413 Request Entity Too Large error.
func EntityTooLargeError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusRequestEntityTooLarge)
	}
	response := NewErrorResponse(message, StatusRequestEntityTooLarge, code)
	response.Send(c)
}

// UnSupportMediaTypeError handles the unsupported media type error by sending an error response to the client.
func UnSupportMediaTypeError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusUnsupportedMediaType)
	}
	response := NewErrorResponse(message, StatusUnsupportedMediaType, code)
	response.Send(c)
}

// InternalServerError represents a 500 Internal Server Error
func InternalServerError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusInternalServerError)
	}
	response := NewErrorResponse(message, StatusInternalServerError, code)
	response.Send(c)

}

// ServiceUnavailable represents a 503 Service Unavailable
func ServiceUnavailable(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = GetReasonPhrase(StatusServiceUnavailable)
	}
	response := NewErrorResponse(message, StatusServiceUnavailable, code)
	response.Send(c)
}
