package service

import (
	"context"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
)

func (t *Todo) List(ctx context.Context) ([]todo.Todo, error) {
	todos, err := t.todoRepo.List(ctx)
	if err != nil {
		return todos, errorsutil.Wrapf(err, "Failed to get todos list", api.ErrCodeInternalDatabase)
	}

	return todos, nil
}
