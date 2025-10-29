package main

import (
	"groupie-tracker/internal/app"
	"groupie-tracker/internal/constants"
	"log"
	"net/http"
)

func main() {
	server, err := app.New(constants.PORT)
	if err != nil {
		log.Printf("failed to initialize app: %v", err)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		})
		port := constants.PORT
		log.Fatal(http.ListenAndServe(port, nil))
		return
	}
	
	log.Printf("listening on %s", server.GetPort())
	log.Fatal(http.ListenAndServe(server.GetPort(), server))
}
