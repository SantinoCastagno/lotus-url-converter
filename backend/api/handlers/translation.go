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
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var translation models.Translation
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
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	rows, err := e.Db.Query("SELECT id, source, destination FROM translations")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var translations []models.Translation
	for rows.Next() {
		var translation models.Translation
		err := rows.Scan(&translation.ID, &translation.Source, &translation.Destionation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		translations = append(translations, translation)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(translations)
}

func (e *Env) GetTranslation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	source := r.URL.Query().Get("source")
	if source == "" {
		http.Error(w, "Missing 'source' parameter", http.StatusBadRequest)
		return
	}

	var translation models.Translation
	err := e.Db.QueryRow("SELECT id, source, destination FROM translations WHERE source = $1", translation.Source).Scan(&translation.ID, &translation.Source, &translation.Destionation)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Translation not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(translation)
}
