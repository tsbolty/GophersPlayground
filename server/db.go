package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tsbolty/GophersPlayground/db/models/todos"
	"github.com/tsbolty/GophersPlayground/db/models/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// You can define a struct to hold your GORM DB connection
type Store struct {
	DB             *gorm.DB
	UserRepository users.UserRepository
	TodoRepository todos.TodoRepository
}

// NewStore creates a new Store with a database connection
func NewStore(db *gorm.DB) *Store {
	return &Store{DB: db}
}

func InitializeDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	log.Printf("Connecting to db: %s\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		migrate(db)
		log.Println("Setting up db connection pool")
		_, err := db.DB()
		if err != nil {
			log.Fatalf("Could not get sql.DB: %v", err)
		}
		log.Println("Connected to db successfully")
		return db, nil
	}

	log.Printf("Could not connect to db: %v", err)
	return nil, err
}

func migrate(db *gorm.DB) error {
	log.Println("Starting AutoMigrate...")
	if err := db.AutoMigrate(&users.User{}, &todos.Todo{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
		return err
	}
	log.Println("AutoMigrate completed")
	return nil
}
