package repository

import (
	"database/sql"

	"github.com/iivkis/vk-chunte/config"
	_ "github.com/lib/pq"
)

type store struct {
	db *sql.DB
}

func newStore() *store {
	db, err := sql.Open("postgres", config.DB_URL)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &store{db: db}
}
