package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Driver de PostgreSQL
)

// var DB *sqlx.DB



func InitDB() (db *sqlx.DB) {
    var err error


    	// URL de conexión
	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
    db, err = sqlx.Connect("postgres", dbURL)
    if err != nil {
        log.Fatalf("Error al conectar con la base de datos: %v", err)
    }
    log.Println("Conexión exitosa a PostgreSQL")
    return db
}



// var DB *sqlx.DB



// func InitDB(dataSourceName string) {
//     var err error
//     DB, err = sqlx.Connect("postgres", dataSourceName)
//     if err != nil {
//         log.Fatalf("Error al conectar con la base de datos: %v", err)
//     }
//     log.Println("Conexión exitosa a PostgreSQL")
// }