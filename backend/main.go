package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/Pallinder/go-randomdata"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"

	"github.com/SantinoCastagno/lotus-url-converter/routes"
)

func connectDB() (*sql.DB, error) {
	connStr := "host=localhost port=5431 user=myuser password=mypassword dbname=mydatabase sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)

		return nil, err
	}

	// Pool of connections configuration
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(8)
	db.SetConnMaxLifetime(time.Minute * 5)

	// Ping to the database
	err = db.Ping()
	if err != nil {
		fmt.Println("Error al hacer ping a la base de datos:", err)
		return nil, err
	}

	log.Info().Msg("Conexión exitosa a la base de datos.")
	return db, nil
}

func main() {

	// library testing
	for i := 0; i < 3; i++ {
		fmt.Println(randomdata.SillyName())
	}

	db, _ := connectDB()
	wrappedMux := routes.SetupRoutes(db)
	http.ListenAndServe(":8090", wrappedMux)

}
