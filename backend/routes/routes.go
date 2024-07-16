package routes

import (
	"database/sql"
	"net/http"

	"github.com/SantinoCastagno/lotus-url-converter/api/handlers"
)

func SetupRoutes(db *sql.DB) {
	env := &handlers.Env{Db: db}

	http.HandleFunc("/get-translation", env.GetTranslation)
	//http.HandleFunc("/post-translation", postTranslation)
}
