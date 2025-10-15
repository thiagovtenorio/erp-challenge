package dbconnectiongo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:1234@localhost:5432/erp")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetConnString() string {
	return "postgres://postgres:1234@localhost:5432/erp"
}
