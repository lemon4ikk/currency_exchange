package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Db *sql.DB
}

func InitDB(path string) (*DB, error) {
	database, err := sql.Open("sqlite3", path)

	return &DB{Db: database}, err
}
