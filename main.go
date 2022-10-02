package main

import (
	"log"
	"net/http"
	"os"
	"server/handlers"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/app/fileHola/", handlers.ReadFile).Methods("GET")
	mux.HandleFunc("/app/template/", handlers.TemplateFile).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3300"
	}

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
