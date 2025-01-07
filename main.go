package main

import (
        "fmt"
        "log"
        "net"
        "os"

        pb "corezn/api/proto"
        "corezn/internal/config"
        "corezn/internal/database"
        "corezn/internal/service"

        "google.golang.org/grpc"
        "google.golang.org/grpc/reflection"
)

const (
        version = "1.0.0"
)

func main() {
        // Load configuration
        cfg := config.LoadConfig()

        // Initialize database
        db, err := database.NewDatabase(cfg.DatabaseURL)
        if err != nil {
                log.Fatalf("Failed to initialize database: %v", err)
                os.Exit(1)
        }

        // Create a gRPC server
        grpcServer := grpc.NewServer()

        // Register our services
        healthCheckServer := service.NewHealthCheckServer(version)
        pb.RegisterHealthCheckServer(grpcServer, healthCheckServer)

        authServer := service.NewAuthServer(db)
        pb.RegisterAuthServer(grpcServer, authServer)

        // Register reflection service on gRPC server
        reflection.Register(grpcServer)

        // Create listener
        addr := fmt.Sprintf(":%s", cfg.ServerPort)
        listener, err := net.Listen("tcp", addr)
        if err != nil {
                log.Fatalf("Failed to listen: %v", err)
                os.Exit(1)
        }

        log.Printf("Starting gRPC server on %s in %s mode", addr, cfg.Environment)
        if err := grpcServer.Serve(listener); err != nil {
                log.Fatalf("Failed to serve: %v", err)
                os.Exit(1)
        }
}