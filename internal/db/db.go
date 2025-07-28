package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {

	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("DATABASE_URL from env:", dbURL)
	if dbURL == "" {
		panic("DATABASE_URL not set in environment")
	}
	var err error
	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to DB: %v", err))
	}

	fmt.Println("Connected to PSQL successfully")
}
