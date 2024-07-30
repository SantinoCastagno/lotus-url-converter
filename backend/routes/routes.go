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
	http.HandleFunc("/generate-translation", env.GenerateTranslation)
	http.HandleFunc("/get-translation", env.GetTranslation)
	http.HandleFunc("/get-translations", env.GetTranslations)
	wrappedMux := handlers.CorsMiddleware(mux)

	return wrappedMux
}
