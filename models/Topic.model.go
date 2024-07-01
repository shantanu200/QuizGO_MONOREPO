package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Topic struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Subject     string             `bson:"subject" json:"subject"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Resources   []string           `bson:"resources" json:"resources"`
	Notes       []string           `bson:"notes" json:"notes"`
}
