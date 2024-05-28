package response

import (
	"github.com/gin-gonic/gin"
)

// SuccessResponse represents a structured success response
type SuccessResponse struct {
	Message          string      `json:"message"`
	Status           int         `json:"status"`
	ReasonStatusCode string      `json:"reason_status_code"`
	Option           interface{} `json:"option,omitempty"`
	Metadata         interface{} `json:"metadata,omitempty"`
}

// NewSuccessResponse creates a new SuccessResponse
func NewSuccessResponse(message string, statusCode int, reasonStatusCode string, option interface{}, metadata interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message:          message,
		Status:           statusCode,
		ReasonStatusCode: reasonStatusCode,
		Option:           option,
		Metadata:         metadata,
	}
}

// Send sends the success response to the client
func (sr *SuccessResponse) Send(c *gin.Context) {
	c.JSON(sr.Status, sr)
}

// Ok represents a 200 OK success response
func Ok(c *gin.Context, message string, metadata interface{}) {
	if message == "" {
		message = GetReasonPhrase(StatusOK)
	}
	response := NewSuccessResponse(message, StatusOK, GetReasonPhrase(StatusOK), nil, metadata)
	response.Send(c)
}

// Created represents a 201 Created success response
func Created(c *gin.Context, message string, metadata interface{}) {
	if message == "" {
		message = GetReasonPhrase(StatusCreated)
	}
	response := NewSuccessResponse(message, StatusCreated, GetReasonPhrase(StatusCreated), nil, metadata)
	response.Send(c)
}
