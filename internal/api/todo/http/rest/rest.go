package rest

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
)

type TodoService interface {
	Create(ctx context.Context, params todo.CreateParams) (todo.Todo, error)
	List(ctx context.Context) ([]todo.Todo, error)
	Find(ctx context.Context, id int64) (todo.Todo, error)
	Update(ctx context.Context, id int64, params todo.UpdateParams) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type TodoHandler struct {
	validate *validator.Validate
	todoSvc  TodoService
}

func NewTodoHandler(todoSvc TodoService) *TodoHandler {
	return &TodoHandler{
		todoSvc: todoSvc,
	}
}

func (h *TodoHandler) Register(r *gin.Engine) {
	r.GET("/todos", h.list)
	r.POST("/todos", h.create)
	r.GET("/todos/:id", h.find)
	r.PUT("/todos/:id", h.update)
	r.DELETE("/todos/:id", h.delete)
}

type Todo struct {
	ID     int64  `json:"id,omitempty"`
	Task   string `json:"task,omitempty"`
	IsDone bool   `json:"is_done,omitempty"`
}
