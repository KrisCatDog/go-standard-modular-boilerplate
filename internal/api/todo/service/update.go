package service

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
)

func (t *Todo) Update(ctx context.Context, id int64, params todo.UpdateParams) (int64, error) {
	updatedID, err := t.todoRepo.Update(ctx, id, params)
	if err != nil {
		if pgxscan.NotFound(err) {
			return updatedID, errorsutil.Wrapf(err, "Todo doesn't exist", api.ErrCodeNotFound)
		}

		return updatedID, errorsutil.Wrapf(err, "Failed to update a todo", api.ErrCodeInternalDatabase)
	}

	return updatedID, nil
}
