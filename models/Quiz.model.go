package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quiz struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title         string               `bson:"title" json:"title"`
	Description   string               `bson:"description" json:"description"`
	InitDate      string               `bson:"initDate" json:"initDate"`
	DueDate       string               `bson:"dueDate" json:"dueDate"`
	Duration      int64                `bson:"duration" json:"duration"`
	MaxScore      int64                `bson:"maxScore" json:"maxScore"`
	TotalQuestion int64                `bson:"totalQuestion" json:"totalQuestion"`
	Questions     []primitive.ObjectID `bson:"questions" json:"questions"`
	QuestionModels []Question          `json:"questionModels"`
}
