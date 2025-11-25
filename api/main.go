package main

import (
    "github.com/gin-gonic/gin"
    "api/middleware"
	"api/routes"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
    r := gin.Default()
	// Register security middleware
	r.Use(scp("localhost:8080"))
    routes.RegisterRoutes(r)
	// Start server on port 8000 (default: 8080)
    r.Run(":8000")
}
