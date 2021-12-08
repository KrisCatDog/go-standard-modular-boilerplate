package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
)

// TodoRepository defines the datastore contract to be implemented.
type TodoRepository interface {
	Create(ctx context.Context, params todo.CreateParams) (todo.Todo, error)
	List(ctx context.Context) ([]todo.Todo, error)
	Find(ctx context.Context, id int64) (todo.Todo, error)
	Update(ctx context.Context, id int64, params todo.UpdateParams) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

// Todo defines the application service with the required dependencies.
type Todo struct {
	todoRepo TodoRepository
	log      *zap.Logger
}

// NewTodo returns an instance of Todo service.
func NewTodo(logger *zap.Logger, todoRepo TodoRepository) *Todo {
	return &Todo{
		todoRepo: todoRepo,
		log:      logger,
	}
}
