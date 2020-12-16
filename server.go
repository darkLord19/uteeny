package main

import (
	"database/sql"
	"net/http"
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
