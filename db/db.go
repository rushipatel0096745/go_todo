package db

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/lib/pq"
)

func Connect() *sql.DB {
    db, err := sql.Open("postgres", os.Getenv("DB_URL"))
    if err != nil {
        panic(fmt.Sprintf("failed to connect to database: %v", err))
    }

    if err := db.Ping(); err != nil {
        panic(fmt.Sprintf("database unreachable: %v", err))
    }

    fmt.Println("connected to supabase postgres")
    return db
}