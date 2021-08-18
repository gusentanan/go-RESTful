package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConnection struct {
	DB_name     string
	DB_username string
	DB_password string

	DB_protocol string
	DB_address  string

	DB *sql.DB
}

type DatabaseChan struct {
	DB  *sql.DB
	err error
}

func Connection(d DatabaseConnection) chan DatabaseChan {
	c := make(chan DatabaseChan)

	go func() {
		defer close(c)
		dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", d.DB_username, d.DB_password, d.DB_protocol, d.DB_address, d.DB_name)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			c <- DatabaseChan{DB: nil, err: err}
		}

		c <- DatabaseChan{DB: db, err: err}
	}()

	return c
}

func (database DatabaseConnection) InitializeDatabase(ctx context.Context) (*sql.DB, error) {
	c := Connection(database)
	for {
		select {
		case <-ctx.Done():
			err := fmt.Errorf("timeout connecting to database : %w", ctx.Err())
			return nil, err
		case db := <-c:
			return db.DB, db.err
		}
	}

}
