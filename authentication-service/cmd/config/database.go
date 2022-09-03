package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// counts for count time response database
var counts int64

// Function open connection to database
func openDB(dsn string) (*sql.DB, error) {
	// Open connection
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Ping db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// If no error, return db
	return db, nil
}

func SetupDB() *sql.DB {
	// Get DSN fron environment variable
	dsn := os.Getenv("DSN")

	for {
		// Open conection to db using function openDB
		conn, err := openDB(dsn)

		if err != nil {
			// Print log
			log.Println("Postgres not yet ready ...")
			// Increments var counts
			counts++
		} else {
			// Print log
			log.Println("Connected to Postgres!")
			// return connection
			return conn
		}

		// If var count is greater that 1-
		if counts > 10 {
			// Print log error from connection
			log.Printf("Database connection error: %s\n", err)
			return nil
		}

		// Print log for waiting two second each trying connection again
		log.Println("Backing off for two seconds ...")
		// Time sleep at 2 secondf
		time.Sleep(2 * time.Second)
		continue
	}
}
