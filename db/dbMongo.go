package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
}

const uri = "mongodb+srv://DrakoMaster:xnkKnbYuGbGVegXP@drako-db.fguhd.mongodb.net/?retryWrites=true&w=majority"

var Client *mongo.Client
var CollTest *mongo.Collection
var DbTest *mongo.Database
var GlobalCtx context.Context
var CancelFun context.CancelFunc

func MountClientMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	clientCon, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	} else {
		GlobalCtx = ctx
		Client = clientCon
		CancelFun = cancel
	}
	defer Client.Disconnect(ctx)

	DbTest = Client.Database("Go_Test")
	CollTest = DbTest.Collection("Primera_Coll")

	title := "Drako"

	var result bson.M
	err = CollTest.FindOne(context.TODO(), bson.D{{"nombre", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the Name %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

}

func CreateRegistro() {
	podcast := Podcast{
		Title:  "No existe el POST",
		Author: "Drako",
	}
	insertResult, err := CollTest.InsertOne(GlobalCtx, &podcast)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println(insertResult.InsertedID)

}
