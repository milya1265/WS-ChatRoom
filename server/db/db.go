package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) { //					postgres://dmilya:qwerty@localhost:5432/BlogDB?sslmode=disable
	db, err := sql.Open("postgres", "postgres://dmilyano:qwerty@localhost:5432/chatroomdb?sslmode=disable")
	if err != nil {
		log.Println(err.Error())
		return nil, err

	}

	if err = db.Ping(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {

	return d.db
}
