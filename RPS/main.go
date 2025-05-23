package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	// Crear un enrutador
	router := http.NewServeMux()
	// Manejador para servir los archivos estaticos
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	// Configurar rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/New", handlers.NewGame)
	router.HandleFunc("/Game", handlers.Game)
	router.HandleFunc("/Play", handlers.Play)
	router.HandleFunc("/About", handlers.About)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
