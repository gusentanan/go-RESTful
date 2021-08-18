package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"
)

func InitializeMainDatabase(ctx context.Context) *sql.DB {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(10000*time.Millisecond))
	defer cancel()
	DBcon := DatabaseConnection{
		DB_name:     os.Getenv("DB_NAME"),
		DB_username: os.Getenv("DB_USERNAME"),
		DB_password: os.Getenv("DB_PASSWORD"),
		DB_protocol: os.Getenv("DB_PROTOCOL"),
		DB_address:  os.Getenv("DB_ADDRESS"),
	}
	db, err := DBcon.InitializeDatabase(ctx)
	if err != nil {
		log.Fatalf("an error occured on connecting to database : %s", err.Error())
	} else {
		log.Fatalf("connection to database success")
	}

	return db
}
