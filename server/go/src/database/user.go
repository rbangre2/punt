package database

import (
	"time"
	"database/sql"
	"punt/models"
)

func CreateUser(db *sql.DB, u models.User) error {
	query := `INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)`
    _, err := db.Exec(query, u.Username, u.Email, u.Password, u.CreatedAt)
    return err
}

func GetUserByEmail(db *sql.DB, email string) (models.User, error) {
    var u models.User
    var createdAt string // Use a string to initially scan the date

    query := `SELECT id, username, email, password, created_at FROM users WHERE email = ?`
    err := db.QueryRow(query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &createdAt)
    if err != nil {
        return u, err
    }

    // Parse the date string into the time.Time type
    // Use the format "2006-01-02 15:04:05" to match MySQL's datetime format
    u.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
    if err != nil {
        return u, err
    }

    return u, nil
}