package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type Env struct {
	Db *sql.DB
}

func getTranslation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		source := r.URL.Query().Get("source")
		fmt.Fprintf(w, "Source: %s!", source)

		// TODO: Get data from the db
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
