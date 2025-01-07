package models

import "time"

type Todo struct {
        ID          string    `json:"id"`
        Title       string    `json:"title" validate:"required,min=1,max=100"`
        Description string    `json:"description" validate:"max=500"`
        Completed   bool      `json:"completed"`
        CreatedAt   time.Time `json:"created_at"`
        UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoRequest struct {
        Title       string `json:"title" validate:"required,min=1,max=100"`
        Description string `json:"description" validate:"max=500"`
}

type UpdateTodoRequest struct {
        Title       *string `json:"title,omitempty" validate:"omitempty,min=1,max=100"`
        Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
        Completed   *bool   `json:"completed,omitempty"`
}

type ErrorResponse struct {
        Error string `json:"error"`
}