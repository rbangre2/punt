package handlers 

import (
    "encoding/json"
    "net/http"
    "punt/database"
    "punt/models"

    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)


func Login(db *sql.DB) http.HandlerFunc {
	var err error
	
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	var jwtSecret := os.Getenv("JWT_SECRET")
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