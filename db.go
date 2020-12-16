package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// connect to postgresql database
func connect(uname, password, dburl, dbport, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", uname, password, dburl, dbport, dbname)
	return sql.Open("postgres", connStr)
}

func createTablesIfNotExist(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
		hash varchar(8) PRIMARY KEY
		original varchar(1024) NOT NULL
		timestamp bigint
	)`
	_, err := db.Exec(query)
	return err
}

func createIndexIfNotExist(db *sql.DB) error {
	query := `CREATE INDEX url_hash ON urls(hash)`
	_, err := db.Exec(query)
	return err
}
