package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	connStr := "postgres://postgres:password@localhost:5432/recordings"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	fmt.Println("Connected to the database!")
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
