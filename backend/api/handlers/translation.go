package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/SantinoCastagno/lotus-url-converter/models"
)

type Env struct {
	Db *sql.DB
}

func (e *Env) GenerateTranslation(w http.ResponseWriter, r *http.Request) {
	var translation models.Translation

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&translation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = e.Db.Exec("INSERT INTO translations (source, destination) VALUES ($1, $2)", translation.Source, "example-example-example")
	if err != nil {
		fmt.Println("Error al insertar datos:", err)
		return
	}

	// generate response and return answer
	response := map[string]string{"msg": "New translation has been added to database."}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (e *Env) CreateTranslation(w http.ResponseWriter, r *http.Request) {
	var _ models.Translation
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	source := r.URL.Query().Get("source")
	fmt.Fprintf(w, "Source: %s!", source)

	//_, err := e.Db.Exec()

	// TODO: Get data from the db
}

// func (e *Env) GetTranslation(w http.ResponseWriter, r *http.Request) {
// 	var result []models.Translation
// 	if r.Method == http.MethodGet {
// 		source := r.URL.Query().Get("source")
// 		fmt.Fprintf(w, "Source: %s!", source)

// 		// TODO: Get data from the db
// 	} else {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 	}
// }
