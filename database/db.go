package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	dbURL := os.Getenv("DB_DSN")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set in .env")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")
	return db
}
