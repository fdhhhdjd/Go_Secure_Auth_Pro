package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	third_party "github.com/fdhhhdjd/Go_Secure_Auth_Pro/third_party/telegram"
	"github.com/gin-gonic/gin"
)

type bodyWriter struct {
	gin.ResponseWriter
	writer io.Writer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

// FormatErrorMessage formats an error message with Markdown for readability.
func FormatErrorMessage(method, path string, duration time.Duration, status int, errorName, formattedRequestBody string) string {
	message := fmt.Sprintf(
		"🚨 *Error!* 🚨\n"+
			"**Request:** `%s`\n"+
			"**Method:** `%s`\n"+
			"**Duration:** `%s`\n"+
			"**Status:** `%d`\n"+
			"**Error Message:** `%s`\n"+
			"**Body:** \n"+
			"```json\n%s\n```",
		path, method, duration, status, errorName, formattedRequestBody)

	return message
}

// RequestLoggingMiddleware is a middleware function that logs information about incoming requests and outgoing responses.
// It captures the request method, path, duration, status code, error name (if any), and request body.
// The captured information is formatted as a Markdown message and sent to a third-party service for logging.
func RequestLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Copy body for request
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			requestBodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))

		method, path := c.Request.Method, c.Request.URL.Path

		// Write again  response body
		responseBodyBuffer := new(bytes.Buffer)
		writer := io.MultiWriter(c.Writer, responseBodyBuffer)
		c.Writer = &bodyWriter{ResponseWriter: c.Writer, writer: writer}

		c.Next()

		duration := time.Since(startTime)
		status := c.Writer.Status()

		responseBodyBytes := responseBodyBuffer.Bytes()
		errorName := ""
		if status >= 500 {
			var responseBody map[string]interface{}
			if err := json.Unmarshal(responseBodyBytes, &responseBody); err == nil {
				if msg, exists := responseBody["message"]; exists {
					errorName = fmt.Sprintf("%v", msg)
				}
			}
		}

		// Convert request body to JSON
		var formattedRequestBody bytes.Buffer
		if err := json.Indent(&formattedRequestBody, requestBodyBytes, "", "  "); err != nil {
			formattedRequestBody.Write(requestBodyBytes)
		}

		// Send message to Telegram
		message := FormatErrorMessage(method, path, duration, status, errorName, formattedRequestBody.String())

		if status >= 500 {
			go third_party.SendTelegramMessage(message, "Markdown", true, false)
		}
	}
}
