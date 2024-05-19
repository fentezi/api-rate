package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDatabase(DbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
