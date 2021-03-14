package flags

import "os"

type mode struct {
	Dev  string
	Prod string
}

var DbCollection = mode{
	Dev:  "images",
	Prod: "images",
}
var DbConnectString = mode{
	Dev:  "mongodb+srv://ngosangns:jikmli@cluster0.oxs6m.mongodb.net/ngosangns?retryWrites=true&w=majority",
	Prod: "mongodb+srv://ngosangns:jikmli@cluster0.oxs6m.mongodb.net/ngosangns?retryWrites=true&w=majority",
}
var DbName = mode{
	Dev:  "ngosangns-myunsplash",
	Prod: "ngosangns-myunsplash",
}

func (varEnv mode) Get() string {
	if os.Getenv("APP_ENV") == "dev" {
		return varEnv.Dev
	}
	if os.Getenv("APP_ENV") == "prod" {
		return varEnv.Prod
	}
	return ""
}
