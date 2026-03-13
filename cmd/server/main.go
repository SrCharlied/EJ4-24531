package main

import (
	"log"
	"net/http"
	"os"

	"go-http/interno/handlers"
	"go-http/interno/storage"
)

func main() {

	// segun esto se config el puerto donde se va a ejecutar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "24531"
	}

	// se inicializa el almacenamiento
	store, err := storage.NewJSONStore("data/games.json")
	if err != nil {
		log.Fatal("Error cargando la data:", err)
	}

	log.Println("Games loaded:", len(store.Games))

	//se crea el router
	mux := http.NewServeMux()

	//se registran los handlers
	gameHandler := handlers.NewGameHandler(store)

	mux.HandleFunc("/api/games", gameHandler.HandleGames)
	mux.HandleFunc("/api/games/", gameHandler.HandleGameByID)

	//endpoint pa probar...
	mux.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	//se inicializa el server
	log.Println("Servidor escuchando en el puerto ", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
