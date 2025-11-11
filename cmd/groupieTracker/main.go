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
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		})
		log.Fatal(http.ListenAndServe(constants.PORT, nil))
		return
	}

	log.Printf("Listening on %s", server.GetPort())
	log.Fatal(http.ListenAndServe(server.GetPort(), server))
}
