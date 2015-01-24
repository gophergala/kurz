package storage

/**
TODO use a CLI flag to clear the database on init
*/

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Storage struct {
	Name string
	DB   *sql.DB
}

var Service Storage = Storage{}

func (s *Storage) Open(dbName string) error {
	var err error
	s.DB, err = sql.Open("mysql", dbName)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (s *Storage) Close() {
	s.DB.Close()
}
