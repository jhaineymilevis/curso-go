package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func ListContacts(db *sql.DB) {

	query :=
		"SELECT id, name, email, phone FROM contact"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("\nLISTA de contactos:")
	fmt.Println("-----------------------------")

	for rows.Next() {
		contact := models.Contact{}

		var valueEmail sql.NullString //para columnas que son nullable
		var valuePhone sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &valuePhone)

		if err != nil {
			log.Fatal(err)
		}
		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin email"
		}

		if valuePhone.Valid {
			contact.Phone = valuePhone.String
		} else {
			contact.Phone = "Sin telefono"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("-----------------------------")
	}

}

func GetContactById(db *sql.DB, contactID int) {
	query := "SELECT id, name, email, phone from contact where id =?"
	row := db.QueryRow(query, contactID)

	contact := models.Contact{}

	var valueEmail sql.NullString
	var valuePhone sql.NullString

	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &valuePhone)

	if err != nil {

		if err == sql.ErrNoRows {
			log.Fatalf("no se encontro ningtun contacto con el id %d", contactID)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "no email"
	}

	if valuePhone.Valid {
		contact.Phone = valuePhone.String
	} else {
		contact.Phone = "no phone"
	}

	fmt.Println("\nImprimiendo contacto")
	fmt.Println("--------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n", contact.Id, contact.Name, contact.Email, contact.Phone)

	fmt.Println("--------------------")

}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?,?,?)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("nuevo contacto creado")
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact set name =?, email =?, phone =? where id = ?"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("clontact actualizado")
}

func DeleteContact(db *sql.DB, contactId int) {

	query := "delete from contact where id =?"

	_, err := db.Exec(query, contactId)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("contacto eliminado")

}
