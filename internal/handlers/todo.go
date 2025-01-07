package handlers

import (
        "encoding/json"
        "net/http"
        "strings"
        "sync"
        "time"
        "todo-service/internal/models"

        "github.com/go-playground/validator/v10"
        "github.com/google/uuid"
)

var (
        todos     = make(map[string]models.Todo)
        mu        sync.RWMutex
        validate  = validator.New()
)

func RegisterRoutes(mux *http.ServeMux) {
        mux.HandleFunc("/health", handleHealth)
        mux.HandleFunc("/todos", handleTodos)
        mux.HandleFunc("/todos/", handleTodo)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
                return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
                "status": "healthy",
                "time":   time.Now().Format(time.RFC3339),
        })
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
                getTodos(w, r)
        case http.MethodPost:
                createTodo(w, r)
        default:
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
        id := strings.TrimPrefix(r.URL.Path, "/todos/")
        if id == "" {
                http.Error(w, "Invalid todo ID", http.StatusBadRequest)
                return
        }

        switch r.Method {
        case http.MethodGet:
                getTodo(w, r, id)
        case http.MethodPut:
                updateTodo(w, r, id)
        case http.MethodDelete:
                deleteTodo(w, r, id)
        default:
                http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
}

func getTodos(w http.ResponseWriter, r *http.Request) {
        mu.RLock()
        todoList := make([]models.Todo, 0, len(todos))
        for _, todo := range todos {
                todoList = append(todoList, todo)
        }
        mu.RUnlock()

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(todoList)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
        var req models.CreateTodoRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid request body"})
                return
        }

        if err := validate.Struct(req); err != nil {
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(models.ErrorResponse{Error: err.Error()})
                return
        }

        todo := models.Todo{
                ID:          uuid.New().String(),
                Title:       req.Title,
                Description: req.Description,
                Completed:   false,
                CreatedAt:   time.Now(),
                UpdatedAt:   time.Now(),
        }

        mu.Lock()
        todos[todo.ID] = todo
        mu.Unlock()

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(todo)
}

func getTodo(w http.ResponseWriter, r *http.Request, id string) {
        mu.RLock()
        todo, exists := todos[id]
        mu.RUnlock()

        if !exists {
                http.Error(w, "Todo not found", http.StatusNotFound)
                return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request, id string) {
        var req models.UpdateTodoRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
        }

        mu.Lock()
        defer mu.Unlock()

        todo, exists := todos[id]
        if !exists {
                http.Error(w, "Todo not found", http.StatusNotFound)
                return
        }

        if req.Title != nil {
                todo.Title = *req.Title
        }
        if req.Description != nil {
                todo.Description = *req.Description
        }
        if req.Completed != nil {
                todo.Completed = *req.Completed
        }
        todo.UpdatedAt = time.Now()

        todos[id] = todo

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request, id string) {
        mu.Lock()
        _, exists := todos[id]
        if !exists {
                mu.Unlock()
                http.Error(w, "Todo not found", http.StatusNotFound)
                return
        }

        delete(todos, id)
        mu.Unlock()

        w.WriteHeader(http.StatusNoContent)
}