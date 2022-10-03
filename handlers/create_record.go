package handlers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateRecord(collection *mongo.Collection,
	ctx context.Context,
	data map[string]interface{}) (map[string]interface{}, error) {

	req, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, err
	}

	insertedId := req.InsertedID

	res := map[string]interface{}{
		"status":  "ok",
		"message": "Nuevo Registro",
		"data":    insertedId,
	}
	// res := map[string]interface{}{
	// 	"status": "ok",
	// 	"message": "Nuevo Registro",
	// 	"data": map[string]interface{}{
	// 		"insertedId": insertedId,
	// 	},
	// }

	return res, nil
}
