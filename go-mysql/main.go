package main

import (
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"

	_ "github.com/go-sql-driver/mysql" // Importar el driver de MySQL de forma indirecta con el _ al inicio
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}
	defer db.Close() // Asegurarse de cerrar la conexión al final

	//handlers.ListContacts(db)
	handlers.GetContactById(db, 2)

	newContact := models.Contact{
		Id:    1,
		Name:  "Miguel",
		Email: "miguel@gmail.com",
		Phone: "2222",
	}

	//handlers.CreateContact(db, newContact)
	handlers.UpdateContact(db, newContact)
	handlers.DeleteContact(db, 1)

}
