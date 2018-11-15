package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Todo contains title, description, status
type Todo struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Status      string        `bson:"status" json:"status"`
}
