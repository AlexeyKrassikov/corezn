# CoreZN gRPC Service

A gRPC-based microservice written in Go.

## Features

- gRPC API
- Health check service
- Configuration through environment variables
- Clean code organization
- Protocol Buffers for service definitions
- gRPC reflection enabled for easy testing

## Services

### Health Check Service
Provides basic health check functionality with version and timestamp information.

```protobuf
service HealthCheck {
  rpc Check(HealthCheckRequest) returns (HealthCheckResponse) {}
}

message HealthCheckRequest {}

message HealthCheckResponse {
  string status = 1;
  string version = 2;
  int64 timestamp = 3;
}
```

## Configuration

The service can be configured using environment variables:

- `SERVER_PORT` - Port to run the gRPC server on (default: "8080")
- `ENVIRONMENT` - Environment name (default: "development")

## Development

### Prerequisites

- Go 1.21 or later
- Protocol Buffers compiler (protoc)
- Go Protocol Buffers plugin
- Go gRPC plugin

### Installing Prerequisites

```bash
# Install Protocol Buffers compiler
apt-get install -y protobuf-compiler

# Install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Generating Protocol Buffers Code

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/proto/healthcheck.proto
```

### Running the Service

```bash
# Run with default configuration
go run main.go

# Run with custom configuration
export SERVER_PORT=8000 ENVIRONMENT=production
go run main.go
```

## Testing with gRPC CLI

You can use tools like `grpcurl` to test the service:

```bash
# List all available services (using reflection)
grpcurl -plaintext localhost:8080 list

# Get service description
grpcurl -plaintext localhost:8080 describe api.HealthCheck

# Call the health check service
grpcurl -plaintext localhost:8080 api.HealthCheck/Check
```