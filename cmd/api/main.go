
package main

import (
	"bootcamp_v1/internal/server"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
	done <- true
}

func main() {
	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize the server
	srv := server.NewServer()

	// Log server startup
	fmt.Printf("🚀 Server is running on http://localhost:%s\n", port)

	// Channel to signal when shutdown is complete
	done := make(chan bool, 1)

	// Start graceful shutdown in a separate goroutine
	go gracefulShutdown(srv, done)

	// Start server
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("❌ HTTP server error: %s", err)
	}

	// Wait for shutdown to complete
	<-done
	log.Println("✅ Graceful shutdown complete.")
}

