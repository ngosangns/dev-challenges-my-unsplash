package main

import (
	"net/http"

	"github.com/ngosangns/devchallenges-my-unsplash-api/api"
)

func main() {
	http.HandleFunc("/api/get", api.Get)
	http.HandleFunc("/api/search", api.Search)
	http.HandleFunc("/api/delete", api.Delete)
	http.HandleFunc("/api/create", api.Create)
	http.ListenAndServe(":8080", nil)
}
