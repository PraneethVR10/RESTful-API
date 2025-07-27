package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDatabaseIfNotExists() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	// Change connection to default `postgres` DB
	bootstrapURL := strings.Replace(dbURL, "/students", "/postgres", 1)

	bootstrapPool, err := pgxpool.New(context.Background(), bootstrapURL)
	if err != nil {
		log.Fatalf("Failed to connect to bootstrap DB: %v", err)
	}
	defer bootstrapPool.Close()

	// Create students DB if not exists
	_, err = bootstrapPool.Exec(context.Background(), "CREATE DATABASE students")
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		log.Fatalf("Failed to create 'students' DB: %v", err)
	}
	fmt.Println("Created or verified 'students' database")
}

// CreateStudentsTable creates the students table
func CreateStudentsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS students (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		admission_num INT NOT NULL
	);`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("Table 'students' created or already exists.")
}
