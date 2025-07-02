package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbUser, dbPassword, dbHost, dbPort, dbName)

	//Abrir la conexion a base de datos

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err

	}

	if err := db.Ping(); err != nil {
		return nil, err

	}
	fmt.Println("Conexion exitosa a la base de datos MySQL")
	return db, nil
}
