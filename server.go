package main

import (
	"database/sql"
	"net/http"
	"os"
)

type env struct {
	db *sql.DB
}

func shorten(e *env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func elongate(e *env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func routes(m *http.ServeMux, e *env) {
	m.Handle("/shorten", shorten(e))
	m.Handle("/elongate", elongate(e))
}

func main() {
	db, err := connect(os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_URL"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	e := env{db: db}
	mux := http.NewServeMux()
	routes(mux, &e)
	http.ListenAndServe(":9090", mux)
}
