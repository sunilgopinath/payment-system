package main

import (
	"context"
	"fmt"

	"github.com/sunilgopinath/payment-system/pkg/db"
)

func main() {
	conn, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Connected to PostgreSQL with pgx!")
	conn.Close(context.Background())
}
