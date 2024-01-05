package main 

import (
    "fmt"
    "log"
    "net/http"
 
    _ "github.com/go-sql-driver/mysql"
    "punt/database"
    "punt/middleware"
    "punt/handlers"
)

func main() {
	db, err := database.Connect()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    defer db.Close()

    http.Handle("/register", middleware.EnableCORS(handlers.RegisterHandler(db)))
    http.Handle("/login", middleware.EnableCORS(handlers.LoginHandler(db)))
    
    fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}