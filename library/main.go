package main

import "library/book"

func main() {

	//Create a instance of Book object
	var myBook = book.NewBook("Mody Dick", "Herman Melville", 300)

	// create the instance of Book without constructor
	/*var myBook = book.Book{
		Title:  "Mody Dick",
		Author: "Herman Melville",
		Pages:  300,
	}*/

	//Using the PrintInfo Book's method
	myBook.PrintInfo()
}
