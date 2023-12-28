package database 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName))
	if err != nil {
		return nil, err
	}
	return db, nil
}
