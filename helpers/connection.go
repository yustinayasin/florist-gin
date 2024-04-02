package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	//connection
	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbName, dbUser, dbPass)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
