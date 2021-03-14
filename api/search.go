package api

import (
	"errors"
	"net/http"

	"github.com/ngosangns/devchallenges-my-unsplash-api/database"
	"github.com/ngosangns/devchallenges-my-unsplash-api/flags"
	"google.golang.org/api/iterator"
)

func Search(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	// Get URL param "id"
	keys, ok := r.URL.Query()["query"]
	if !ok || len(keys[0]) < 1 {
		printErr(w, errors.New("Query mustn't be null"), "")
		return
	}
	query := keys[0]

	// Connect database
	client, ctx, err := database.Connect()
	defer client.Close()
	if err != nil {
		printErr(w, err, "Error while connecting to database")
		return
	}

	// Get records by query
	results := make([]*map[string]interface{}, 0)
	q := client.Collection(flags.DbCollection.Get()).Where("label", ">=", query).Where("label", ">=", query+"\uf8ff")
	docs := q.Documents(ctx)

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
