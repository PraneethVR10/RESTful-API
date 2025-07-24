package db

import (
	"context"
	"fmt"
)

// Run this function ONCE to create the table
func CreateStudentsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS students (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		admission_num INT NOT NULL
	);`

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		panic(fmt.Sprintf("Failed to create table: %v", err))
	}

	fmt.Println("Table 'students' created or already exists.")
}
