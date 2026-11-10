package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func NewConnectionDB() (*sql.DB, error) {

	//Load .env Varibles
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env files")
	}

	DB_NAME := os.Getenv("DB_NAME")
	HOST := os.Getenv("HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	SSL_MODE := os.Getenv("SSL_MODE")
	DB_PORT := os.Getenv("DB_PORT")

	//establishing our database connection

	dsn := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s?sslmode=%s`, DB_USER, DB_PASSWORD, HOST, DB_PORT, DB_NAME, SSL_MODE)

	//Open db connection
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	//verify db connection
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err

}
