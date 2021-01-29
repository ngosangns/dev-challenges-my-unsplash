package handler

import (
	"context"
	"log"
	"net/http"

	"my-unsplash/database"

	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const password string = "123"

func Delete(c echo.Context) error {
	// Get value from request
	id := strings.TrimSpace(c.Param("id"))
	pwd := strings.TrimSpace(c.Request().Header.Get("Password"))

	// Validate
	if id == "" {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"status":  false,
			"message": "Id mustn't be null",
		})
	}

	// Check password
	if pwd != password {
		return c.JSON(http.StatusNotImplemented, map[string]interface{}{
			"status":  false,
			"message": "Wrong password",
		})
	}

	// Connect database
	client, cancel, err := database.Connect()
	if err != nil {
		return err
	}
	defer cancel()
	collection := client.Database(database.DB_NAME).Collection(database.DB_COLLECTION)

	// Remove record
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = collection.DeleteMany(context.TODO(), bson.M{"_id": idPrimitive})
	if err != nil {
		log.Fatal(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})
}
