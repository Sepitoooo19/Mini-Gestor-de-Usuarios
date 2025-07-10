package models

type Example struct {
    ID    string `json:"id" bson:"_id,omitempty"`
    Title string `json:"title" bson:"title"`
}