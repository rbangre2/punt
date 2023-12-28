package models 

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Claims struct for JWT
type Claims struct {
    UserID int `json:"userid"`
    jwt.StandardClaims
}

type User struct {
	ID        int       `json:"id"`
    Username  string    `json:"username"` 
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
}