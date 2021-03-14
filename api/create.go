package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ngosangns/devchallenges-my-unsplash-api/database"
	"github.com/ngosangns/devchallenges-my-unsplash-api/flags"

	"github.com/ngosangns/devchallenges-my-unsplash-api/models"

	"regexp"
)

// Handler
func Create(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	// Get value from request
	label := strings.TrimSpace(r.FormValue("label"))
	link := strings.TrimSpace(r.FormValue("link"))

	// Validate
	if label == "" && link == "" {
		printErr(w, errors.New("The fields can't be null"), "")
		return
	}
	labelPattern := regexp.MustCompile(`^[0-9a-zA-Z_]+$`)
	linkPattern := regexp.MustCompile(`^https:\/\/images.unsplash.com\/`)
	if !labelPattern.MatchString(label) || !linkPattern.MatchString(link) {
		printErr(w, errors.New("The fields don't match the format"), "")
		return
	}

	// Check file type of link
	res, err := http.Get(link)
	if err != nil {
		printErr(w, err, "Error")
		return
	}
	defer res.Body.Close()
	fileType := res.Header.Get("Content-Type")
	if strings.Split(fileType, "/")[0] != "image" {
		printErr(w, errors.New("Isn't an image link"), "")
		return
	}

	// Connect database
	client, ctx, err := database.Connect()
	defer client.Close()
	if err != nil {
		printErr(w, err, "Error while connecting to database")
		return
	}

	// Add record
	ash := models.Image{
		Label: label,
		Link:  link,
	}
	ref, _, err := client.Collection(flags.DbCollection.Get()).Add(ctx, ash)
	if err != nil {
		printErr(w, err, "Error")
		return
	}

	// Response
	printRes(w, map[string]interface{}{
		"_id": ref.ID,
		"data": models.Image{
			Label: label,
			Link:  link,
		},
	})

}
