package api

import (
	"errors"
	"net/http"

	"github.com/ngosangns/devchallenges-my-unsplash-api/database"
	"github.com/ngosangns/devchallenges-my-unsplash-api/flags"

	"strings"
)

const password string = "123"

func Delete(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	// Get URL param "id"
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		printErr(w, errors.New("Id mustn't be null"), "")
		return
	}
	id := keys[0]

	// Get Password header
	pwd := strings.TrimSpace(r.Header.Get("Password"))
	if pwd != password {
		printErr(w, errors.New("Wrong password"), "")
		return
	}

	// Connect database
	client, ctx, err := database.Connect()
	defer client.Close()
	if err != nil {
		printErr(w, err, "Error while connecting to database")
		return
	}

	// Remove record
	ref := client.Collection(flags.DbCollection.Get()).Doc(id)
	// Check record exist
	_, err = ref.Get(ctx)
	if err != nil {
		printErr(w, err, "Error")
		return
	}
	// Remove
	_, err = ref.Delete(ctx)
	if err != nil {
		printErr(w, err, "Error")
		return
	}

	printRes(w, "Deleted successfully")
}
