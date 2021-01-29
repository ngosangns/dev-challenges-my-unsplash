package handler

import (
	"my-unsplash/database"
	"my-unsplash/model"
	"net/http"

	"github.com/labstack/echo/v4"

	"strings"

	"context"

	"go.mongodb.org/mongo-driver/bson"

	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Search(c echo.Context) error {
	// Get value from request
	query := strings.TrimSpace(c.QueryParam("query"))

	// Validate
	if query == "" {
		return c.JSON(http.StatusNotImplemented, model.Response{
			false,
			"Query mustn't be null",
		})
	}

	// Connect database
	client, cancel, err := database.Connect()
	if err != nil {
		return err
	}
	defer cancel()
	collection := client.Database(database.DB_NAME).Collection(database.DB_COLLECTION)

	// Get records by query
	results := make([]*map[string]interface{}, 0)
	cur, err := collection.Find(context.TODO(), bson.D{{
		"label",
		bson.D{{
			"$in",
			bson.A{query},
		}},
	}})
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
