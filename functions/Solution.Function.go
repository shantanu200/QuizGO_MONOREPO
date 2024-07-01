package functions

import (
	"context"
	"fmt"
	"quizGo/configs"
	"quizGo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var SolutionCollection *mongo.Collection = configs.GetCollection(configs.DB, "Solution")

func GetSolution(ctx context.Context, userId string, boardId string) (*models.SolutionModel, error) {
	_user, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	_board, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"boardId": _board, "userId": _user}

	var solutionBoard *models.SolutionModel
	err = SolutionCollection.FindOne(ctx, filter).Decode(&solutionBoard)
	if err != nil {
		return nil, err
	}

	return solutionBoard, nil
}

func CreateSolution(ctx context.Context, userId string, doc map[string]interface{}) (bool, error) {
	_, findErr := GetSolution(ctx, userId, doc["boardId"].(string))

	questionKeySearch := fmt.Sprint("userSolution.", doc["questionKey"])

	_user, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return false, err
	}

	_boardId, err := primitive.ObjectIDFromHex(doc["boardId"].(string))
	if err != nil {
		return false, err
	}

	doc["userId"] = _user
	doc["boardId"] = _boardId

	if findErr != nil {
		_, err = SolutionCollection.InsertOne(ctx, doc)
		if err != nil {
			return false, err
		}

		return true, nil
	} else {
		filter := bson.M{"userId": _user, "boardId": _boardId}
		update := bson.M{"$set": bson.M{questionKeySearch: doc["userSolution"].(map[string]interface{})[doc["questionKey"].(string)]}}

		_, err := SolutionCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func GetAllSolutionForUser(ctx context.Context, userId string, page int64, limit int64) (*models.SolutionModel, error) {
	objUser, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"userId": objUser}

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip((page - 1) * limit)

	cursor, err := SolutionCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var solutions *models.SolutionModel
	if err = cursor.All(ctx, &solutions); err != nil {
		return nil, err
	}

	return solutions, nil
}
