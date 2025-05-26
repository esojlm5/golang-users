package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect(user, password, dbname, host string, port int) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
