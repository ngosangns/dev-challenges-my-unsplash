package models

type Image struct {
	Label string `firestore:"label" json:"label"`
	Link  string `firestore:"link" json:"link"`
}
