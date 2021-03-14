package api

import (
	"net/http"

	"github.com/ngosangns/devchallenges-my-unsplash-api/database"
	"github.com/ngosangns/devchallenges-my-unsplash-api/flags"

	"google.golang.org/api/iterator"
)

func Get(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	// Connect database
	client, ctx, err := database.Connect()
	defer client.Close()
	if err != nil {
		printErr(w, err, "Error while connecting to database")
		return
	}

	// Get all records
	results := make([]*map[string]interface{}, 0)
	collection := client.Collection(flags.DbCollection.Get())
	docs := collection.Documents(ctx)
	defer docs.Stop()
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			printErr(w, err, "Error")
			return
		}
		results = append(results, &map[string]interface{}{
			"_id":  doc.Ref.ID,
			"data": doc.Data(),
		})
	}
	printRes(w, results)
}
