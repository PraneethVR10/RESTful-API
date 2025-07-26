// db/db.go
package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("DATABASE_URL not set in environment")
	}

	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to DB: %v", err))
	}

	fmt.Println("Connected to PSQL successfully")
}
