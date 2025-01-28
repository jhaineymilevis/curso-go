package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	/*----------------------------------------------------------------/
	  	/*
	      Polimorfismo
	*/
	/*----------------------------------------------------------------*/
	//Create a instance of Book object
	var mbBook = book.NewBook("Mody Dick", "Herman Melville", 300, 1090)
	var thgBook = book.NewBook("The hunger games", "Herman Melville", 300, 2024)

	// create the instance of Book without constructor
	/*var myBook = book.Book{
		Title:  "Mody Dick",
		Author: "Herman Melville",
		Pages:  300,
	}*/

	//Using the PrintInfo Book's method
	//mbBook.PrintInfo()
	//thgBook.PrintInfo()

	// Set Year of Mody Dick book
	mbBook.SetYear(2022)
	fmt.Printf("Mody Dick's year: %d\n", mbBook.GetYear())

	textBook := book.NewTextBook(
		"The Catcher in the Rye",
		"J.D. Salinger",
		260,
		"J.D. Salinger",
		"Moderno",
		2025,
	)

	//textBook.PrintInfo()

	book.Print(mbBook)
	book.Print(textBook)
	book.Print(thgBook)

	/*----------------------------------------------------------------/
	  	/*
	      Interfaces
	*/
	/*----------------------------------------------------------------*/

	myDog := animal.Dog{Name: "Max"}
	myCat := animal.Cat{Name: "Chipi"}

	animal.MakeSound(&myDog)
	animal.MakeSound(&myCat)

	animals := []animal.Animal{
		&myCat, &myCat, &animal.Cat{Name: "Kiki"}, &animal.Cat{Name: "Kira"},
	}

	for _, animal := range animals {
		animal.Sound()
	}

}
