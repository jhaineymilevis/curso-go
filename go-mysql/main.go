package main

import (
	"go-mysql/database"

	_ "github.com/go-sql-driver/mysql" // Importar el driver de MySQL de forma indirecta con el _ al inicio
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}
	defer db.Close() // Asegurarse de cerrar la conexión al final

}
