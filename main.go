package main

import (
        "fmt"
        "log"
        "net/http"
        "todo-service/internal/config"
        "todo-service/internal/handlers"
        "todo-service/internal/middleware"
)

func main() {
        // Load configuration
        cfg := config.LoadConfig()

        // Initialize router
        mux := http.NewServeMux()

        // Register routes
        handlers.RegisterRoutes(mux)

        // Create a middleware chain
        handler := middleware.Recovery(middleware.Logger(mux))

        // Start server
        addr := fmt.Sprintf(":%s", cfg.ServerPort)
        log.Printf("Starting server on %s in %s mode", addr, cfg.Environment)
        
        server := &http.Server{
                Addr:    addr,
                Handler: handler,
        }

        if err := server.ListenAndServe(); err != nil {
                log.Fatal(err)
        }
}