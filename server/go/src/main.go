package main 

import (
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "punt/database"
)

func main() {
	db, err := database.Connect()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    defer db.Close()

    // Ping the database to test the connection
    err = db.Ping()
    if err != nil {
        log.Fatalf("Could not ping the database: %v", err)
    }
    fmt.Println("Successfully connected to the database.")

    user, err := database.GetUserByEmail(db, "newuser@example.com")
    if err != nil {
        log.Fatalf("Error getting user by email: %v", err)
    }

    // Print the details of the user retrieved.
    fmt.Printf("User found: %+v\n", user)
}