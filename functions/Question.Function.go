package functions

import (
	"context"
	"fmt"
	"log"
	"quizGo/configs"
	"quizGo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var questionCollection *mongo.Collection = configs.GetCollection(configs.DB, "questions")

func CreateQuestion(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error) {
	result, err := questionCollection.InsertOne(ctx, doc)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func InsertMultipleQuestion(ctx context.Context, doc []interface{}) (*mongo.InsertManyResult, error) {
	result, err := questionCollection.InsertMany(ctx, doc)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetQuestions(ctx context.Context, page int64, limit int64, query string) ([]models.Question, error) {
	filter := bson.M{}

	if query != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"text": bson.M{
						"$regex": primitive.Regex{Pattern: query, Options: "i"},
					},
				},
				{
					"description": bson.M{
						"$regex": primitive.Regex{Pattern: query, Options: "i"},
					},
				},
			},
		}
	}

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip((page - 1) * limit)

	var questions []models.Question

	cursor, err := questionCollection.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &questions); err != nil {
		return nil, err
	}

	return questions, nil
}

func SingleQuestion(ctx context.Context, questionId string) (*models.Question, error) {
	_id, err := primitive.ObjectIDFromHex(questionId)

	if err != nil {
		return nil, err
	}

	var question *models.Question

	err = questionCollection.FindOne(ctx, bson.M{"_id": _id}).Decode(&question)

	if err != nil {
		return nil, err
	}

	return question, nil

}

func UpdateQuestion(ctx context.Context, questionId string, doc map[string]interface{}) (*mongo.UpdateResult, error) {
	objQuestion, _ := primitive.ObjectIDFromHex(questionId)

	fmt.Println(doc)

	if topics, ok := doc["topics"].([]interface{}); ok {
		var objectIds []primitive.ObjectID
		if len(topics) == 0 {
			fmt.Println("Len is zero")
			doc["topics"] = []primitive.ObjectID{}
		} else {
			for _, str := range topics {
				objId, err := primitive.ObjectIDFromHex(str.(string))
				if err != nil {
					log.Fatal(err)
				}
				objectIds = append(objectIds, objId)
			}
			doc["topics"] = objectIds
		}
	}

	update := bson.D{
		{
			Key:   "$set",
			Value: doc,
		},
	}

	result, err := questionCollection.UpdateOne(ctx, bson.M{"_id": objQuestion}, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteQuestion(ctx context.Context, questionId string) (*mongo.DeleteResult, error) {
	objQuestion, _ := primitive.ObjectIDFromHex(questionId)

	result, err := questionCollection.DeleteOne(ctx, bson.M{"_id": objQuestion})

	if err != nil {
		return nil, err
	}

	return result, err
}

func GetTopicWiseQuestion(ctx context.Context, topicId string, page int64, limit int64, query string) ([]models.Question, error) {
	objTopic, _ := primitive.ObjectIDFromHex(topicId)

	var questions []models.Question

	filter := bson.M{"topics": objTopic}

	if query != "" {
		filter = bson.M{
			"topics": objTopic,
			"$or": []bson.M{
				{
					"text": bson.M{
						"$regex": primitive.Regex{Pattern: query, Options: "i"},
					},
				},
				{
					"description": bson.M{
						"$regex": primitive.Regex{Pattern: query, Options: "i"},
					},
				},
			},
		}
	}

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip((page - 1) * limit)

	cursor, err := questionCollection.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &questions); err != nil {
		return nil, err
	}

	return questions, nil
}
