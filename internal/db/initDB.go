package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var (
	conn     *pgx.Conn
	initOnce sync.Once
)

func InitDB() *pgx.Conn {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("No database URL provided")
	}

	initOnce.Do(func() {
		var err error
		conn, err = pgx.Connect(context.Background(), dbURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
	})

	return conn
}
