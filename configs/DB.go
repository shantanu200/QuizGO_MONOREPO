package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"quizGo/constants"
)

func ConnectDatabase() *mongo.Client {
	ctx := context.TODO()

	uri := constants.EnvConstant("DATABASEURI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database is connected to service!")

	return client
}

var DB *mongo.Client = ConnectDatabase()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("quizGo").Collection(collectionName)

	if collectionName == "students" {
		collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
	}

	return collection
}
