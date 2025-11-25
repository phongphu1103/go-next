package middleware

import (
    "github.com/gin-gonic/gin"
	"net/http"
)

// Security-Content-Policy returns a Gin middleware handler
func SecurityHeadersMiddleware(expectedHost string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Basic host header validation
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}

		// Not allow browsers to open the website in an iframe of other websites
		c.Header("X-Frame-Options", "sameorigin")
		// Allow to load resources from known places only
		c.Header("Content-Security-Policy",
			"default-src 'self'; "+
				"connect-src *; "+
				"font-src *; "+
				"script-src-elem * 'unsafe-inline'; "+
				"img-src * data:; "+
				"style-src * 'unsafe-inline';")
		// Prevent from XSS attack
		c.Header("X-XSS-Protection", "1; mode=block")
		// Force to use HTTPS
		c.Header("Strict-Transport-Security", "max-age=86400; includeSubDomains")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy",
			"geolocation=(), midi=(), sync-xhr=(), microphone=(), camera=(), " +
				"magnetometer=(), gyroscope=(), fullscreen=(self), payment=()")

		c.Next()
	}
}
