package postgresql

import (
	"context"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/db"
)

// Update updates an existing Todo record, returns deleted ID.
func (t *Todo) Update(ctx context.Context, id int64, params todo.UpdateParams) (int64, error) {
	updatedID, err := t.q.UpdateTodo(ctx, db.UpdateTodoParams{
		ID:     id,
		Task:   params.Task,
		IsDone: params.IsDone,
	})
	if err != nil {
		return updatedID, err
	}

	return updatedID, nil
}
