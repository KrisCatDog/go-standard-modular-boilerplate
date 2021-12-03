package postgresql

import (
	"context"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/api/todo"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

func (t *Todo) List(ctx context.Context) ([]todo.Todo, error) {
	var todos []todo.Todo

	sql, _, err := sq.Select("id", "task", "is_done").From("todos").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return todos, err
	}

	if err := pgxscan.Select(ctx, t.conn, &todos, sql); err != nil {
		return todos, err
	}

	return todos, nil
}
