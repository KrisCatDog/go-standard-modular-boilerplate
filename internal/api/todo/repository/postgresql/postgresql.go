package postgresql

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/db"
)

type Todo struct {
	conn *pgxpool.Pool
	q    *db.Queries
}

func NewTodo(conn *pgxpool.Pool) *Todo {
	return &Todo{
		conn: conn,
		q:    db.New(conn),
	}
}
