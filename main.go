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

	// HANDLERS APP
	mux.HandleFunc("/app/fileHola/", handlers.ReadFile).Methods("GET")
	mux.HandleFunc("/app/template/", handlers.TemplateFile).Methods("GET")

	// HANDLERS API
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUserId).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.EditUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3300"
	}

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
