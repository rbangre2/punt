// middleware/auth.go
package middleware

import (
    "net/http"
    "strings"
    "os"

    "github.com/dgrijalva/jwt-go"
    "punt/models"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract the token from the Authorization header
        tokenString := r.Header.Get("Authorization")
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        claims := &models.Claims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Token is valid; proceed with the request
        next.ServeHTTP(w, r)
    })
}
