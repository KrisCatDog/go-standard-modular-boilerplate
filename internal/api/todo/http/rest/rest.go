package rest

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
)

// TodoService defines a contract for the handlers to be implemented.
type TodoService interface {
	Create(ctx context.Context, params todo.CreateParams) (todo.Todo, error)
	List(ctx context.Context) ([]todo.Todo, error)
	Find(ctx context.Context, id int64) (todo.Todo, error)
	Update(ctx context.Context, id int64, params todo.UpdateParams) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

// TodoHandler defines a handler with the required dependencies.
type TodoHandler struct {
	validate *validator.Validate
	todoSvc  TodoService
}

// NewTodoHandler returns an instance of TodoHandler.
func NewTodoHandler(validate *validator.Validate, todoSvc TodoService) *TodoHandler {
	return &TodoHandler{
		validate: validate,
		todoSvc:  todoSvc,
	}
}

// Register registers the HTTP REST handlers route.
func (h *TodoHandler) Register(r *gin.Engine) {
	r.GET("/todos", h.list)
	r.POST("/todos", h.create)
	r.GET("/todos/:id", h.find)
	r.PUT("/todos/:id", h.update)
	r.DELETE("/todos/:id", h.delete)
}

// Todo defines an entity for the HTTP REST layer.
type Todo struct {
	ID     int64  `json:"id,omitempty"`
	Task   string `json:"task,omitempty"`
	IsDone bool   `json:"is_done,omitempty"`
}
