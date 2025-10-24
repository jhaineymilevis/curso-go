package handlers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "Piedra, Papel o Tijera",
		Message: "¡Bienvenido al juego!",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error al ejecutar la plantilla", http.StatusInternalServerError)
		return
	}
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/new_game.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
func Game(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/game.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/play.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
