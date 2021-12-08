package service

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/pkg/errorsutil"
)

// Delete removes the existing Todo record with its business logic.
func (t *Todo) Delete(ctx context.Context, id int64) (int64, error) {
	deletedID, err := t.todoRepo.Delete(ctx, id)
	if err != nil {
		if pgxscan.NotFound(err) {
			return deletedID, errorsutil.Wrapf(err, "Todo doesn't exist", api.ErrCodeNotFound)
		}

		return deletedID, errorsutil.Wrapf(err, "Failed to delete a todo", api.ErrCodeInternalDatabase)
	}

	return deletedID, nil
}
