package main

import (
    "context"
    "log"
    "os/signal"
	"syscall"
	"time"
    "github.com/gin-gonic/gin"
    "net/http"
    
    "backend/database"
    "backend/providers"
    "backend/http/middleware"
    "backend/routes"
)

func main() {
    app := providers.NewApp()
    // Register providers
    app.Register(providers.DatabaseServiceProvider{})
    // Boot providers (loads config + connect DB)
    app.Boot()

    // Create a Gin router with default middleware (logger and recovery)
    r := gin.Default()
    r.Use(middleware.RequestLogger())
    // Register security middleware
    r.Use(scp("localhost:8080"))
    routes.RegisterRoutes(r)
    // Start server on port 8000 (default: 8080)
    // r.Run(":8000")
    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    // Start server in goroutine
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server crashed: %v", err)
        }
    }()

    // Create a context that listens for SIGINT / SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
    // Block until signal received
	<-ctx.Done()
	log.Println("Shutdown signal received")

    // Gracefully stop receiving new requests
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Forced shutdown due to error: %v", err)
	}

	// Disconnect from DB
	database.Disconnect()

	log.Println("Server exited properly")
}
