package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	econder := json.NewEncoder(file)
	err = econder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil

}

func main() {

	var contacts []Contact
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Printf("Error loading contacts")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("====== GESTOR DE CONTACTOS ====== \n",
			"1. Agregar contacto \n",
			"2. Mostrar contactos \n",
			"3. Salir \n",
			"Seleccione una opcion: ")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion")
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Print("Ingrese el nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Ingrese el email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Ingrese el telefono: ")
			c.Phone, _ = reader.ReadString('\n')
			contacts = append(contacts, c)
			err = saveContactToFile(contacts)
			if err != nil {
				fmt.Println("Error al guardar los contactos")
			}

		case 2:
			fmt.Println("================================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s - Email: %s - Telefono: %s\n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("================================================")

		case 3:
			return
		default:
			fmt.Println("Opcion invalida")

		}
	}

}
