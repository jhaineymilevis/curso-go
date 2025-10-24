package main

import (
	"log"
	"net/http"
	"rps/handlers"
)

func main() {
	//enrutador
	router := http.NewServeMux()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	//w.Write([]byte("Hola, mundo!"))
	// 	fmt.Fprintln(w, "Hola mundo")
	// })

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
