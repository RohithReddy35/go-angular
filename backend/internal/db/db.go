package db

import (
    "database/sql"
    "log"

    "github.com/RohithReddy35/go-angular/internal/config"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    connStr := config.GetDBConnectionString()
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    if err := DB.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
    log.Println("Database connection established.")
}
