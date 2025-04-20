package db

import (
	"database/sql"
	"log"
)

var (
    PostgresDB *sql.DB
)

func InitPostgres(connStr string) {
    var err error
    PostgresDB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to PostgreSQL: %v", err)
    }
    if err = PostgresDB.Ping(); err != nil {
        log.Fatalf("PostgreSQL ping failed: %v", err)
    }
    log.Println("Connected to PostgreSQL")
}