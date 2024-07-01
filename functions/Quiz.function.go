package functions

import (
	"context"
	"fmt"
	"log"
	"quizGo/cache"
	"quizGo/configs"
	"quizGo/models"
	"quizGo/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var quizCollection *mongo.Collection = configs.GetCollection(configs.DB, "quiz")

func CreateQuiz(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error) {
	result, err := quizCollection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	err = cache.UpdateDocumentCountCache(ctx, "totalQuizzes", 1)
	if err != nil {
		count, err := quizCollection.CountDocuments(ctx, bson.M{})
		if err != nil {
			return nil, err
		}

		err = cache.SetCache(ctx, "totalQuizzes", count)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func GetAllQuizzes(ctx context.Context, page int64, limit int64, query string) (types.FilterPagination, error) {
	var quizzes []models.Quiz
	filter := bson.M{}

	var totalDocuments int

	if exists := cache.CheckKeyCache(ctx, "totalQuizzes"); exists {
		err := cache.GetCache(ctx, "totalQuizzes", &totalDocuments)
		if err != nil {
			fmt.Println("Exiting from cache call")
			return types.FilterPagination{}, err
		}
	} else {
		count, err := quizCollection.CountDocuments(ctx, bson.M{})

		fmt.Println("Database Call => Total Count = ", count)

		if err != nil {
			return types.FilterPagination{}, err
		}

		err = cache.SetCache(ctx, "totalQuizzes", count)
		if err != nil {
			return types.FilterPagination{}, err
		}
	}

	fmt.Println(totalDocuments)

	if query != "" {
		filter = bson.M{
			"$or": []bson.M{
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

	cursor, err := quizCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return types.FilterPagination{}, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &quizzes); err != nil {
		return types.FilterPagination{}, err
	}

	if query != "" {
		totalDocuments = len(quizzes)
	}

	return types.FilterPagination{
		TotalResults: totalDocuments,
		PageSize:     limit,
		PageNumber:   page,
		Results:      quizzes,
	}, nil
}

func UpdateQuiz(ctx context.Context, quizId string, doc map[string]interface{}) (*mongo.UpdateResult, error) {
	objQuiz, _ := primitive.ObjectIDFromHex(quizId)

	if topics, ok := doc["questions"].([]interface{}); ok {
		var objectIds []primitive.ObjectID
		if len(topics) == 0 {
			fmt.Println("Question Array length is zero")
			doc["questions"] = []primitive.ObjectID{}
		} else {
			for _, str := range topics {
				objId, err := primitive.ObjectIDFromHex(str.(string))
				if err != nil {
					log.Fatal(err)
				}
				objectIds = append(objectIds, objId)
			}
			doc["questions"] = objectIds
		}
	}

	update := bson.D{
		{
			Key:   "$set",
			Value: doc,
		},
	}

	result, err := quizCollection.UpdateOne(ctx, bson.M{"_id": objQuiz}, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteQuiz(ctx context.Context, quizId string) (*mongo.DeleteResult, error) {
	objQuiz, _ := primitive.ObjectIDFromHex(quizId)

	result, err := quizCollection.DeleteOne(ctx, bson.M{"_id": objQuiz})
	if err != nil {
		return nil, err
	}

	err = cache.UpdateDocumentCountCache(ctx, "totalQuizzes", -1)
	if err != nil {
		count, err := quizCollection.CountDocuments(ctx, bson.M{})
		if err != nil {
			return nil, err
		}

		err = cache.SetCache(ctx, "totalQuizzes", count)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func GetSingleQuiz(ctx context.Context, quizId string) (models.Quiz, error) {
	CacheKey := "quiz-" + quizId
	var quiz models.Quiz

	if exists := cache.CheckKeyCache(ctx, CacheKey); exists {
		fmt.Println("Pulling Data from cache")
		err := cache.GetCache(ctx, CacheKey, &quiz)
		if err != nil {
			return models.Quiz{}, err
		}
	} else {
		fmt.Println("Pulling Quiz Details from database")

		objQuiz, _ := primitive.ObjectIDFromHex(quizId)

		err := quizCollection.FindOne(ctx, bson.M{"_id": objQuiz}).Decode(&quiz)
		if err != nil {
			return models.Quiz{}, err
		}

		var question []models.Question

		filter := bson.M{"_id": bson.M{"$in": quiz.Questions}}
		cursor, err := questionCollection.Find(ctx, filter)
		if err != nil {
			return models.Quiz{}, err
		}

		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &question); err != nil {
			return models.Quiz{}, err
		}

		quiz.QuestionModels = question

		err = cache.SetCache(ctx, CacheKey, quiz)
		if err != nil {
			return models.Quiz{}, err
		}
	}

	return quiz, nil
}
