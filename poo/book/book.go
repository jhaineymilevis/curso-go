package book

import "fmt"

/*
Book object (the equivalent to Class in common POO languages)
*/
type Book struct {

	//Book public properties. Can be accesed out of package
	Title  string
	Author string
	Pages  int
	//Book private properties. Only can be accesed in the package
	year int
}

/* Simulate Book Constructor  */
func NewBook(title string, author string, pages int, year int) *Book {

	//it returns a pointer of book object
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
		year:   year,
	}
}

// Getters and Setters of Privated Properties
func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) GetYear() int {

	return b.year
}

/*
A method of Book object,
It have a receptor, the receptor is a pointer to the Book object, it simulate the class method of common POO languages
*/
func (b *Book) PrintInfo() {

	fmt.Printf("Title: %s\nAuthor: %s\nPages: % d\n", b.Title, b.Author, b.Pages)
}

/* A child struct of Book, Simulating "Herencia" */
type TextBook struct {
	Book
	edittorial string
	level      string
}

/* constructor of TextBook */
func NewTextBook(title, author string, pages int, editorial, level string, year int) *TextBook {

	return &TextBook{
		Book:       *NewBook(title, author, pages, 0), //book object
		edittorial: editorial,
		level:      level,
	}
}

/*
A method of TextBook object,
It have a receptor, the receptor is a pointer to the TextBook object, it simulate the class method of common POO languages
*/
func (b *TextBook) PrintInfo() {

	fmt.Printf("Title: %s\nAuthor: %s\nPages: % d\nYear: %d\n", b.Title, b.Author, b.Pages, b.year)
}
