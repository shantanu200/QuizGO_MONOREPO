package functions

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"quizGo/configs"
	"quizGo/middlewares"
	"quizGo/models"
	"quizGo/utils"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "students")

func CreateUser(ctx context.Context, doc interface{}) (string, error) {

	result, err := userCollection.InsertOne(ctx, doc)

	if err != nil {
		if strings.Contains(err.Error(), "E11000") {
			return "", errors.New("email already present")
		}
		return "", err
	}

	token, err := middlewares.CreateAuthToken(result.InsertedID.(primitive.ObjectID).Hex())

	if err != nil {
		return "", nil
	}

	return token, nil
}

func GetUser(ctx context.Context, userId string) (*models.Users, error) {
	objUser, _ := primitive.ObjectIDFromHex(userId)

	var user *models.Users

	err := userCollection.FindOne(ctx, bson.M{"_id": objUser}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	var user *models.Users

	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUser(ctx context.Context, email string, password string) (string, error) {
	user, err := GetUserByEmail(ctx, email)

	if err != nil {
		return "", errors.New("user not found")
	}

	if user.Password != password {
		return "", errors.New("invalid Password")
	}

	token, err := middlewares.CreateAuthToken(user.ID.Hex())

	if err != nil {
		return "", errors.New("unable to create token")
	}

	return token, nil

}

func UpdateUser(ctx context.Context, userId string, doc map[string]interface{}) (*mongo.UpdateResult, error) {
	objUser, _ := primitive.ObjectIDFromHex(userId)

	update := bson.D{
		{
			Key:   "$set",
			Value: doc,
		},
	}

	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objUser}, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteUser(ctx context.Context, userId string) (*mongo.DeleteResult, error) {
	objUser, _ := primitive.ObjectIDFromHex(userId)

	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objUser})

	if err != nil {
		return nil, err
	}

	return result, err
}

func InsertMultipleUsers(ctx context.Context) (*mongo.InsertManyResult, error) {
	users := utils.CreateRandomFakeUser()

	results, err := userCollection.InsertMany(ctx, users)

	if err != nil {
		return nil, err
	}

	return results, nil
}
