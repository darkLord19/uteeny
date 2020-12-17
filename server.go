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
		if r.Method != "GET" {
			http.Error(w, r.Method+" is not allowed on /shorten", http.StatusMethodNotAllowed)
			return
		}
		o, exist := r.URL.Query()["url"]
		if !exist || len(o) == 0 {
			http.Error(w, "please provide url to shorten", http.StatusBadRequest)
			return
		}
		u := url{original: o[0], hash: ""}
		u.calculateHash()
		err := store(e.db, u)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		data := u.toJSON()
		if string(data) == "" {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func elongate(e *env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		u, err := hashLookup(e.db, path[1:])
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, u.original, http.StatusSeeOther)
	}
}

func routes(m *http.ServeMux, e *env) {
	m.Handle("/shorten", shorten(e))
	m.Handle("/", elongate(e))
}

func main() {
	db, err := connect(os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_URL"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	err = createTablesIfNotExist(db)
	if err != nil {
		panic(err)
	}
	createIndexIfNotExist(db)
	e := env{db: db}
	mux := http.NewServeMux()
	routes(mux, &e)
	http.ListenAndServe(":9090", mux)
}
