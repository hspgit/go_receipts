package server

import (
	"github.com/gin-gonic/gin"
)

// ErrorMiddleware ErrorResponse standardizes API error responses
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Only handle errors if there are any
		if len(c.Errors) > 0 {
			c.JSON(-1, gin.H{
				"errors": c.Errors.Errors(),
			})
		}
	}
}
