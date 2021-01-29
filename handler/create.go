package handler

import (
	"context"
	"my-unsplash/database"
	"my-unsplash/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"log"
	"regexp"
)

// Handler
func Create(c echo.Context) error {
	// Get value from request
	label := strings.TrimSpace(c.FormValue("label"))
	link := strings.TrimSpace(c.FormValue("link"))

	// Validate
	if label == "" && link == "" {
		return c.JSON(http.StatusNotImplemented, model.Response{
			Status:  false,
			Message: "The fields can't be null",
		})
	}
	labelPattern := regexp.MustCompile(`^[0-9a-zA-Z_]+$`)
	linkPattern := regexp.MustCompile(`^https:\/\/images.unsplash.com\/`)
	if !labelPattern.MatchString(label) || !linkPattern.MatchString(link) {
		return c.JSON(http.StatusNotImplemented, model.Response{
			Status:  false,
			Message: "The fields don't match the format",
		})
	}

	// Check file type of link
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer res.Body.Close()
	fileType := res.Header.Get("Content-Type")
	if strings.Split(fileType, "/")[0] != "image" {
		return c.JSON(http.StatusNotImplemented, model.Response{
			Status:  false,
			Message: "Isn't an image link",
		})
	}

	// Connect database
	client, cancel, err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer cancel()
	collection := client.Database(database.DB_NAME).Collection(database.DB_COLLECTION)

	// Add record
	ash := model.Image{
		Label: label,
		Link:  link,
	}
	queryResponse, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"_id": queryResponse.InsertedID.(primitive.ObjectID).Hex(),
		"data": model.Image{
			Label: label,
			Link:  link,
		},
	})

}
