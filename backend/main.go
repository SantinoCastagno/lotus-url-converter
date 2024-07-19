package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"

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

	fmt.Println("Conexión exitosa a la base de datos.")
	return db, nil
}

func main() {
	db, _ := connectDB()
	routes.SetupRoutes(db)
	http.ListenAndServe(":8090", nil)
}
