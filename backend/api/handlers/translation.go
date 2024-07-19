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

func (e *Env) GetTranslations(w http.ResponseWriter, r *http.Request) {
	var translations []models.Translation
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	//TODO: Continue from here

	// generate response and return answer
	response := map[string]string{"msg": "New translation has been added to database."}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
