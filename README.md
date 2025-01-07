# CoreZN Todo Service

A simple, production-ready TODO microservice written in Go.

## Features

- RESTful API for managing TODO items
- Input validation
- Logging middleware
- Error handling middleware
- Health check endpoint
- Configuration through environment variables
- Clean code organization

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /todos` - List all todos
- `POST /todos` - Create a new todo
- `GET /todos/{id}` - Get a specific todo
- `PUT /todos/{id}` - Update a todo
- `DELETE /todos/{id}` - Delete a todo

## Configuration

The service can be configured using environment variables:

- `SERVER_PORT` - Port to run the server on (default: "8080")
- `ENVIRONMENT` - Environment name (default: "development")

## Running the Service

```bash
# Run with default configuration
go run main.go

# Run with custom configuration
export SERVER_PORT=8000 ENVIRONMENT=production
go run main.go
```

## Example Usage

### Create a Todo
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries","description":"Get milk and bread"}' \
  http://localhost:8080/todos
```

### List All Todos
```bash
curl http://localhost:8080/todos
```

### Get a Specific Todo
```bash
curl http://localhost:8080/todos/{id}
```

### Update a Todo
```bash
curl -X PUT -H "Content-Type: application/json" \
  -d '{"completed":true}' \
  http://localhost:8080/todos/{id}
```

### Delete a Todo
```bash
curl -X DELETE http://localhost:8080/todos/{id}
```