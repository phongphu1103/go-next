package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger logs all incoming requests with duration, status and path.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf(
			"[GIN] %d | %v | %s %s",
			c.Writer.Status(),
			time.Since(start),
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}