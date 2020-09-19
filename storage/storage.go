package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sqlx.DB
}

func Open(url string) (*DB, error) {
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return &DB{DB: db}, nil
}