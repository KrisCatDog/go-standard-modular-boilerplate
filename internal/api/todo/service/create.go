package service

import (
	"context"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
)

// Create stores a new Todo record with its business logic.
func (t *Todo) Create(ctx context.Context, params todo.CreateParams) (todo.Todo, error) {
	newTodo, err := t.todoRepo.Create(ctx, params)
	if err != nil {
		return todo.Todo{}, errorsutil.Wrapf(err, "Failed to create new todo", api.ErrCodeInternalDatabase)
	}

	return newTodo, nil
}
