package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/SantinoCastagno/lotus-url-converter/models"
	_ "github.com/lib/pq"
)

type Env struct {
	Db *sql.DB
}

func GetTranslation(w http.ResponseWriter, r *http.Request) {
	var result []models.Translation
	if r.Method == http.MethodGet {
		source := r.URL.Query().Get("source")
		fmt.Fprintf(w, "Source: %s!", source)

		// TODO: Get data from the db
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func PostTranslation(w http.ResponseWriter, r *http.Request) {
	var result []models.Translation
	if r.Method == http.MethodPost {
		source := r.URL.Query().Get("source")
		fmt.Fprintf(w, "Source: %s!", source)

		// TODO: Get data from the db
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
