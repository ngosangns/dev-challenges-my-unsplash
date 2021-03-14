package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	models "github.com/ngosangns/devchallenges-my-unsplash-api/models"
)

// Util handler
func Util(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	printErr(w, errors.New("404 Not found"), "")
}

func setHeader(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Headers", "Password")
	}
}

func printRes(w http.ResponseWriter, res interface{}) {
	resJSON, err := json.Marshal(models.Res{
		Status:  true,
		Message: res,
	})
	if err != nil {
		printErr(w, err, "Error")
	}
	w.Write(resJSON)
}

func printErr(w http.ResponseWriter, err error, clientErr string) {
	// Print log
	log.Println(err)
	// Set client message
	if clientErr == "" {
		clientErr = err.Error()
	}
	// Print response
	b, _ := json.Marshal(models.Res{
		Status:  false,
		Message: clientErr,
	})
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
}

func regEx(str string, pattern string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}
