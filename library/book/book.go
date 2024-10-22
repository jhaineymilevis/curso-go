package book

import "fmt"

/*
Book object (the equivalent to Class in common POO languages)
*/
type Book struct {

	//Book properties
	Title  string
	Author string
	Pages  int
}

/* Simulate Book Constructor  */
func NewBook(title string, author string, pages int) *Book {

	//it returns a pointer of book object
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

/*
A method of Book object,
It have a receptor, the receptor is a pointer to the Book object, it simulate the class method of common POO languages
*/
func (b *Book) PrintInfo() {

	fmt.Printf("Title: %s\nAuthor: %s\nPages: % d\n", b.Title, b.Author, b.Pages)
}
