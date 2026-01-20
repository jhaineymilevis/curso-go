package main

import (
	"log"
	"net/http"
	"rps/handlers"
)

func main() {
	//enrutador
	router := http.NewServeMux()

	//manejo de archivos estaticos
	fs := http.FileServer(http.Dir("static"))
	//ruta para accder a lso achivos estaticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	//iniciar el servidor

	port := ":8080"
	log.Printf("Servidor escuchando en el puerto http://localhost%s", port)
	//http.ListenAndServe(port, nil)
	log.Fatal(http.ListenAndServe(port, router))
}
