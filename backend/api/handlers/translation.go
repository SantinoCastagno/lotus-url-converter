package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/SantinoCastagno/lotus-url-converter/models"
)

type Env struct {
	Db *sql.DB
}

func (e *Env) GenerateTranslation(w http.ResponseWriter, r *http.Request) {
	var res models.Translation
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	conn, err := e.Db.Conn(c)

	source := r.URL.Query().Get("source")
	fmt.Fprintf(w, "Source: %s!", source)
	// TODO: Get data from the db

}

func (e *Env) GetTranslation(w http.ResponseWriter, r *http.Request) {
	var result []models.Translation
	if r.Method == http.MethodGet {
		source := r.URL.Query().Get("source")
		fmt.Fprintf(w, "Source: %s!", source)

		// TODO: Get data from the db
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
