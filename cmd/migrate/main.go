package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fluxx1on/finance_transaction_system/internal/config"
	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), config.NewDB().URL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	dbFile, err := os.Open("db.sql")
	if err != nil {
		log.Fatalf("Unable to open migration file %v", err)
	}

	migrationSQL, err := io.ReadAll(dbFile)
	if err != nil {
		log.Fatalf("Unable to read migration file: %v", err)
	}

	_, err = conn.Exec(context.Background(), string(migrationSQL))
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migrations successfully applied.")
}
