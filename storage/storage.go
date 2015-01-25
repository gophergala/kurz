/*
The "storage" package in Kurz provides the data model.

In its current implementation:

  - it is specific to MySQL/MariaDB and compatibles SQL engines.
  - data model and data access are not isolated
  - it assumes the DB has been initialized by importing the data/db-schema.sql file
  - it initializes the connection info based on the "dsn" CLI flag
*/
package storage

/**
TODO use a CLI flag to clear the database on init
TODO use a CLI flag to install the database schema
*/

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	DB  *sql.DB
	DSN string
}

var Service Storage = Storage{}

/*
Open() established the database connection, making sure it is actually available instead of just preparing the connection info.
*/
func (s *Storage) Open() error {
	var err error

	s.DB, err = sql.Open("mysql", s.DSN)
	if err == nil {
		err = s.DB.Ping()
	}

	return err
}

/*
Close() closes the database connection and releases its information.
*/
func (s *Storage) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
	s.DB = nil
}

/*
SetDSN() sets the DSN information for the storage.
*/
func (s *Storage) SetDSN(dsn string) {
	s.DSN = dsn
}

/*
init() initializes the storage information from the command-line flag "dsn".
*/
func init() {
	var dsn = flag.String("dsn", "root:@tcp(localhost:3306)/go_kurz", "some_user:some_pass@tcp(some_host:some_port)/some_db")
	flag.Parse()
	Service.SetDSN(*dsn)
}
