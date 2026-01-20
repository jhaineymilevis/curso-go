package handlers

import (
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "new_game.html", nil)
}
func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "game.html", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "play.html", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tmpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Error al ejecutar la plantilla", http.StatusInternalServerError)
		log.Println("Error al ejecutar la plantilla:", err)
		return
	}
}
