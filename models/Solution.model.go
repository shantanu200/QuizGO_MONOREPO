package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SolutionModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	QuizId       primitive.ObjectID `bson:"quizId" json:"quizId"`
	UserId       primitive.ObjectID `bson:"userId" json:"userId"`
	BoardId      primitive.ObjectID `bson:"boardId" json:"boardId"`
	UserSolution interface{}        `bson:"userSolution" json:"userSolution"`
}
