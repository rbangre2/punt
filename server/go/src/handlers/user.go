package handlers 

import (
    "encoding/json"
    "net/http"
    "punt/database"
    "punt/models"
    "log"
    "os"
    "time"
    
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "database/sql"
	"github.com/joho/godotenv"
)


func LoginHandler(db *sql.DB) http.HandlerFunc {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	jwtSecret := os.Getenv("JWT_SECRET")
	var jwtKey = []byte(jwtSecret)

	return func(w http.ResponseWriter, r *http.Request) {
        var creds models.Credentials // Assuming you have a Credentials struct with Email and Password
        err := json.NewDecoder(r.Body).Decode(&creds)
        if err != nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }

        user, err := database.GetUserByEmail(db, creds.Email)
        if err != nil {
            http.Error(w, "User not found", http.StatusUnauthorized)
            return
        }

        if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
            http.Error(w, "Wrong password", http.StatusUnauthorized)
            return
        }

        // Token creation
        expirationTime := time.Now().Add(30 * time.Minute)
        claims := &models.Claims{
            UserID: user.ID,
            StandardClaims: jwt.StandardClaims{
                ExpiresAt: expirationTime.Unix(),
            },
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        tokenString, err := token.SignedString(jwtKey)
        if err != nil {
            http.Error(w, "Server error", http.StatusInternalServerError)
            return
        }

        http.SetCookie(w, &http.Cookie{
            Name:    "token",
            Value:   tokenString,
            Expires: expirationTime,
        })
    }
}

func RegisterHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user models.User
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            http.Error(w, "Error hashing password", http.StatusInternalServerError)
            return
        }
        user.Password = string(hashedPassword)
        user.CreatedAt = time.Now()

        // Create user in the database
        err = database.CreateUser(db, user)
        if err != nil {
            log.Printf("Error creating user: %v", err) // This will print the specific error
            http.Error(w, "Error creating user", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
    }
}