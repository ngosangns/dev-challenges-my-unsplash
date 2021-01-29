package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"my-unsplash/handler"
)

func main() {
	// Set environment variables
	// os.Setenv("PORT", "80")
	// os.Setenv("DB_CONNECT_STRING", "mongodb://localhost:27017")
	// os.Setenv("DB_NAME", "admin")
	// os.Setenv("DB_COLLECTION", "my-unsplash")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	// Limit size
	e.Use(middleware.BodyLimit("4M"))
	// Disable CROS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// Routes
	e.GET("/", handler.Hello)
	e.GET("/get", handler.Get)
	e.GET("/search", handler.Search)
	e.GET("/delete/:id", handler.Delete)
	e.POST("/create", handler.Create)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
