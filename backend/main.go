package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func connectDB() (*sql.DB, error) {
	connStr := "user=myuser password=mypassword dbname=mydatabase sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:", err)

		return nil, err
	}

	// Pool of connections configuration
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
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
	//db, _ := connectDB()

	http.ListenAndServe(":8090", nil)
}
