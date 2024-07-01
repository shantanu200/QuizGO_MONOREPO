package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuizBoard struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	QuizId         primitive.ObjectID `bson:"quizId" json:"quizId"`
	QuizModel      Quiz               `bson:"quizModel" json:"quizModel"`
	UserId         primitive.ObjectID `bson:"userId" json:"userId"`
	User           Users
	QuizToken      string      `bson:"quizToken" json:"quizToken"`
	IsActive       bool        `bson:"isActive" json:"isActive"`
	IsOver         bool        `bson:"isOver" json:"isOver"`
	QuizAttempt    interface{} `bson:"quizAttempt" json:"quizAttempt"`
	CorrectAnswers int         `bson:"correctAnswers" json:"correctAnswers"`
	WrongAnswers   int         `bson:"wrongAnswers" json:"wrongAnswers"`
	TotalAttempted int         `bson:"totalAttempted" json:"totalAttempted"`
	Accuracy       float64     `bson:"accuracy" json:"accuracy"`
	TotalScore     float64     `bson:"totalScore" json:"totalScore"`
}
