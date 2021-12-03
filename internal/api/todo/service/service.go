package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
)

type TodoRepository interface {
	Create(ctx context.Context, params todo.CreateParams) (todo.Todo, error)
	List(ctx context.Context) ([]todo.Todo, error)
	Find(ctx context.Context, id int64) (todo.Todo, error)
	Update(ctx context.Context, id int64, params todo.UpdateParams) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type Todo struct {
	todoRepo TodoRepository
	log      *zap.Logger
}

func NewTodo(logger *zap.Logger, todoRepo TodoRepository) *Todo {
	return &Todo{
		todoRepo: todoRepo,
		log:      logger,
	}
}
