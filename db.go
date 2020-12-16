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
