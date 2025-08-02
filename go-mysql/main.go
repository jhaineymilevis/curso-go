package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" // Importar el driver de MySQL de forma indirecta con el _ al inicio
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}
	defer db.Close() // Asegurarse de cerrar la conexión al final

	handlers.GetContactById(db, 2)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1: Listar contactos")
		fmt.Println("2: Obtener contacto por id")
		fmt.Println("3: Crear nuevo contacto")
		fmt.Println("4: Actualizar contacto")
		fmt.Println("5: Eliminar contacto")
		fmt.Println("6: Salir")
		fmt.Println("Seleccione una opcion")

		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			handlers.ListContacts(db)

		case 2:
			fmt.Println("Ingresa el id del contacto:")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactById(db, idContact)

		case 3:
			newContact := inoputContactDetails(3)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)

		case 4:
			updateContact :=
				inoputContactDetails(4)
			handlers.UpdateContact(db, updateContact)
			handlers.ListContacts(db)
		case 5:
			fmt.Println("Ingresa el id del contacto que deseas elimar:")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
		case 6:
			fmt.Println("saliendo del programa")
			return

		default:
			fmt.Println("Opcion incorrects")
		}
	}

}

func inoputContactDetails(option int) models.Contact {
	//leer entrada del usuairo usando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		fmt.Println("Ingresa el id del contacto que deseas editar:")
		var idContact int
		fmt.Scanln(&idContact)
		contact.Id = idContact
	}

	fmt.Println("ingresa el nombre del contacto")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Println("ingresa el email del contacto")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Println("ingresa el telefono del contacto")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
