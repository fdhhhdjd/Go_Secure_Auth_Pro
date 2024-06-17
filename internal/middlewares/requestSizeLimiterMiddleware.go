package middlewares

import (
	"bytes"
	"io"
	"net/http"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// RequestSizeLimiter is a middleware function that limits the size of incoming requests.
// It sets the maximum size for the request body and returns an error if the request size exceeds the limit.
// If the request size is within the limit, the function allows the request to proceed to the next middleware or handler.
func RequestSizeLimiter(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap the request body with an io.LimitReader to enforce the max size
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)

		// Read the request body into a buffer to check its size
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			if err == io.EOF {
				response.EntityTooLargeError(c, response.ErrorBodySizeTooLarge)
				return
			}
			response.BadRequestError(c, response.ErrorNotRead)
			return
		}

		// Restore the request body so it can be read by other handlers
		c.Request.Body = io.NopCloser(io.LimitReader(io.MultiReader(io.NopCloser(c.Request.Body), io.NopCloser(bytes.NewReader(body))), maxSize))

		c.Next()
	}
}
