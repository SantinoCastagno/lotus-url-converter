package routes

import (
	"database/sql"
	"net/http"

	"github.com/SantinoCastagno/lotus-url-converter/api/handlers"
)

func SetupRoutes(db *sql.DB) {
	// define a pointer to centralizing connection handling to the database
	env := &handlers.Env{Db: db}

	http.HandleFunc("/generate-translation", env.GenerateTranslation)
	http.HandleFunc("/get-translation", env.GetTranslation)
	http.HandleFunc("/get-translations", env.GetTranslations)
}
