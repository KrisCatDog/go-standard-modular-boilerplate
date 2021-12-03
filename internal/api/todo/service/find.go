package service

import (
	"context"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
	"github.com/georgysavva/scany/pgxscan"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
)

func (t *Todo) Find(ctx context.Context, id int64) (todo.Todo, error) {
	singleTodo, err := t.todoRepo.Find(ctx, id)
	if err != nil {
		if pgxscan.NotFound(err) {
			return singleTodo, errorsutil.Wrapf(err, "Todo doesn't exist", api.ErrCodeNotFound)
		}

		return singleTodo, errorsutil.Wrapf(err, "Failed to find a todo", api.ErrCodeInternalDatabase)
	}

	return singleTodo, nil
}
