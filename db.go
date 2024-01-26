// db.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	// Retrieve PostgreSQL details from environment variables
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// Omitting the fmt.Sprintf for the connection string

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create the "csv_data" table if it doesn't exist
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS csv_data (
      id SERIAL PRIMARY KEY,
      name VARCHAR(255),
      age INTEGER,
      occupation VARCHAR(255)
    )
  `)
	if err != nil {
		return nil, err
	}

	// Log that the connection is being attempted
	log.Printf("Connecting to PostgreSQL...")

	return db, nil
}
