package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

const postgresConn = "postgresql://sunilgopinath@localhost:5432/payments?sslmode=disable"

func ConnectDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), postgresConn)
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}
	return conn, nil
}
