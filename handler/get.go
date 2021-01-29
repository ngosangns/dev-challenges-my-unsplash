package handler

import (
	"context"
	"log"
	"my-unsplash/database"
	"net/http"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"

	"my-unsplash/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Get(c echo.Context) error {
	// Connect database
	client, cancel, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer cancel()
	collection := client.Database(database.DB_NAME).Collection(database.DB_COLLECTION)

	// Get all records
	results := make([]*map[string]interface{}, 0)
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
		return err
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var image model.Image
		var imageId struct {
			Value primitive.ObjectID `bson:"_id"`
		}

		// Decode
		err := cur.Decode(&image)
		if err != nil {
			log.Fatal(err)
			return err
		}
		err = cur.Decode(&imageId)
		if err != nil {
			log.Fatal(err)
			return err
		}

		results = append(results, &map[string]interface{}{
			"_id":  imageId.Value.Hex(),
			"data": image,
		})
	}

	return c.JSON(http.StatusOK, results)
}
