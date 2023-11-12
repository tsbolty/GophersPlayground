package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func connectDB(maxRetries int) (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	log.Println("DATABASE CONNECTION STRING!!!!!:", dbURL)

	var db *sql.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", dbURL)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		log.Printf("Could not connect to db: %v, retrying...", err)
		time.Sleep(2 * time.Second) // Wait for two seconds before the retry
	}
	return nil, fmt.Errorf("after %d attempts, last error: %s", maxRetries, err)
}

func OpenDB() (*sql.DB, error) {
	db, err := connectDB(5)
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to db %s successfully", os.Getenv("DB_NAME"))
	return db, nil
}
