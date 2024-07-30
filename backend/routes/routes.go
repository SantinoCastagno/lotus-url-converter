package routes

import (
	"database/sql"
	"net/http"

	"github.com/SantinoCastagno/lotus-url-converter/api/handlers"
)

func SetupRoutes(db *sql.DB) http.Handler {
	// define a pointer to centralizing connection handling to the database
	env := &handlers.Env{Db: db}
	mux := http.NewServeMux()
	mux.HandleFunc("/generate-translation", env.GenerateTranslation)
	mux.HandleFunc("/get-translation", env.GetTranslation)
	mux.HandleFunc("/get-translations", env.GetTranslations)
	corsMux := handlers.CorsMiddleware(mux)
	return corsMux
}
