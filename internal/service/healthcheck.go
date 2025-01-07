package service

import (
	"context"
	"time"

	pb "corezn/api/proto"
)

type HealthCheckServer struct {
	pb.UnimplementedHealthCheckServer
	version string
}

func NewHealthCheckServer(version string) *HealthCheckServer {
	return &HealthCheckServer{
		version: version,
	}
}

func (s *HealthCheckServer) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status:    "healthy",
		Version:   s.version,
		Timestamp: time.Now().Unix(),
	}, nil
}