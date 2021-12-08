package postgresql

import (
	"context"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
)

// Find returns the Todo by querying using its ID.
func (t *Todo) Find(ctx context.Context, id int64) (todo.Todo, error) {
	result, err := t.q.FindTodo(ctx, id)
	if err != nil {
		return todo.Todo{}, err
	}

	return todo.Todo{
		ID:     result.ID,
		Task:   result.Task,
		IsDone: result.IsDone,
	}, nil
}
