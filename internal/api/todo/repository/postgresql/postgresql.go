package postgresql

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/db"
)

// Todo defines a postgres repository with the required dependencies.
type Todo struct {
	conn *pgxpool.Pool
	q    *db.Queries
}

// NewTodo returns an instance of Todo repository.
func NewTodo(conn *pgxpool.Pool) *Todo {
	return &Todo{
		conn: conn,
		q:    db.New(conn),
	}
}
