package main

import (
	"fmt"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)

}

func loadPage(title string) (*Page, error) {

	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}

func handler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path[1:] // Elimina la barra inicial
	fmt.Fprintf(w, "Hola! me encantan los %s", message)
}

func main() {
	// p1 := &Page{Title: "Test page", Body: []byte("This is a test page")}
	// p1.save()

	// p2, err := loadPage("Test page")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(p2.Body))
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
