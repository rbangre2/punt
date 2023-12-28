package database 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	var err error 
	db, err = sql.Open("mysql", "root:Darkfox72@/punt")
	if err != nil {
		return nil, err
	}
	return db, nil
}
