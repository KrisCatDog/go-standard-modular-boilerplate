package postgresql

import (
	"context"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/db"
)

// Create insert a new Todo record, return the new Todo just created.
func (t *Todo) Create(ctx context.Context, params todo.CreateParams) (todo.Todo, error) {
	result, err := t.q.CreateTodo(ctx, db.CreateTodoParams{
		Task:   params.Task,
		IsDone: params.IsDone,
	})
	if err != nil {
		return todo.Todo{}, err
	}

	return todo.Todo{
		ID:   result.ID,
		Task: result.Task,
	}, nil
}
