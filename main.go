package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/handlers"
	"server/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Uri = "mongodb+srv://DrakoMaster:xnkKnbYuGbGVegXP@drako-db.fguhd.mongodb.net/?retryWrites=true&w=majority"

func requestHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{}
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(Uri))

	if err != nil {
		fmt.Println(err.Error())
	}
	DbTest := client.Database("Go_Test")
	collection := DbTest.Collection("Primera_Coll")
	data := map[string]interface{}{}
	err = json.NewDecoder(req.Body).Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
	}

	switch req.Method {
	case "POST":
		response, err = handlers.CreateRecord(collection, ctx, data)
	case "GET":
		response, err = handlers.GetRecords(collection, ctx)
	case "PUT":
		response, err = handlers.UpdateRecord(collection, ctx, data)
	case "DELETE":
		response, err = handlers.DeleteRecord(collection, ctx, data)
	default:
		utils.BadResponse(w, utils.RespBad{
			Message:    "Metodo no permitido",
			StatusCode: http.StatusMethodNotAllowed,
		})
	}

	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	} else {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")

		if err := enc.Encode(response); err != nil {
			fmt.Println(err.Error())
		}
	}

}

func main() {
	mux := mux.NewRouter()

	// mux.HandleFunc("/api/v1/get/", handlers.CreateUser).Methods("GET")
	mux.HandleFunc("/api/v1/pokes/", handlers.GetPokes).Methods("GET")
	mux.HandleFunc("/api/v1/banksAve/", handlers.ListBanksAve).Methods("POST")
	mux.HandleFunc("/api/v1/products/", requestHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3300"
	}

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
