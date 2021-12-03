package postgresql

import (
	"github.com/KrisCatDog/go-standard-modular-boilerplate/internal/db"
	"github.com/jackc/pgx/v4/pgxpool"
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
