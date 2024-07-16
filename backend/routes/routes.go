package routes

import (
	"database/sql"
	"net/http"
	
)

func SetupRoutes(db *sql.DB) {
	//env := &handlers.Env{Db: db}

	http.HandleFunc("/get-translation", getTranslation)
	//http.HandleFunc("/post-translation", postTranslation)
}
