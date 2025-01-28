package main

import (
	"fmt"
	"log"

	"github.com/jhaineymilevis/greetings"
)

func main() {

	log.SetPrefix("greetings")
	log.SetFlags(0)
	message, error := greetings.Hello("Jhay")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Printf(message)
	names := []string{"John", "Jane", "Bob"}

	messages, err2 := greetings.Hellos(names)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(messages)
}
