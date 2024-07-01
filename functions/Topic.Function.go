package functions

import (
	"context"
	"quizGo/configs"
	"quizGo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var topicCollection *mongo.Collection = configs.GetCollection(configs.DB, "topics")

func CreateTopic(ctx context.Context, doc map[string]interface{}) (*mongo.InsertOneResult, error) {
	result, err := topicCollection.InsertOne(ctx, doc)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllTopics(ctx context.Context, page int64, limit int64, query string) ([]models.Topic, error) {
	var topics []models.Topic

	filter := bson.M{}

	if query != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"subject": bson.M{
						"$regex": primitive.Regex{Pattern: query, Options: "i"},
					},
				},
				{
					"title": bson.M{
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

	cursor, err := topicCollection.Find(ctx, filter, findOptions)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &topics); err != nil {
		return nil, err
	}

	return topics, nil
}

func UpdateTopicDetails(ctx context.Context, topicId string, doc interface{}) (*mongo.UpdateResult, error) {
	objTopicId, _ := primitive.ObjectIDFromHex(topicId)
	update := bson.D{
		{
			Key:   "$set",
			Value: doc,
		},
	}

	filter := bson.M{"_id": objTopicId}

	result, err := topicCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteTopicDetails(ctx context.Context, topicId string) (*mongo.DeleteResult, error) {
	objTopicId, _ := primitive.ObjectIDFromHex(topicId)
	filter := bson.M{"_id": objTopicId}

	result, err := topicCollection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return result, nil

}
