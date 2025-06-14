package main

import (
	"fmt"
	"html/template"
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {

		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderRemplates(w, "view.html", p)
}

func editPageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}

	}

	renderRemplates(w, "edit.html", p)

}

func savePageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	fmt.Println("Title:", title)
	fmt.Println("Body:", body)
	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}

// voy a cargar todas las plantillas para manejarlas en cache
var templates = template.Must(template.ParseFiles("view.html", "edit.html"))

func renderRemplates(w http.ResponseWriter, tmpl string, p *Page) {

	// ya no usamos parseFile sporque ya las cargamos en cache
	// t, err := template.ParseFiles(tmpl)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	err = t.Execute(w, p)

	err := templates.ExecuteTemplate(w, tmpl, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Template rendered successfully")

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
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editPageHandler)
	http.HandleFunc("/save/", savePageHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
