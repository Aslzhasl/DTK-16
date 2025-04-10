package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to open DB: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("DB unreachable: ", err)
	}
	log.Println("Connected to PostgreSQL")
	return db
}
