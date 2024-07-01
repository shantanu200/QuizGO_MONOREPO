package functions

import (
	"context"
	"fmt"
	"quizGo/configs"
	"quizGo/middlewares"
	"quizGo/models"
	"quizGo/utils"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QuizAttempt struct {
	Answer string `json:"answer"`
	Score  int    `json:"score"`
}

var quizBoardCollection *mongo.Collection = configs.GetCollection(configs.DB, "quizBoard")

func GetQuizBoard(ctx context.Context, userId string, quizId string) (*models.QuizBoard, error) {
	var quizBoard *models.QuizBoard

	objQuiz, err := primitive.ObjectIDFromHex(quizId)
	if err != nil {
		return nil, err
	}

	objUser, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	err = quizBoardCollection.FindOne(ctx, bson.M{"quizId": objQuiz, "userId": objUser}).Decode(&quizBoard)
	if err != nil {
		return nil, err
	}

	return quizBoard, nil
}

func UpdateQuizBoard(ctx context.Context, boardId string, doc map[string]interface{}) (*mongo.UpdateResult, error) {
	objBoard, _ := primitive.ObjectIDFromHex(boardId)

	totalScore := 0.0
	correctAnswers := 0
	wrongAnswers := 0
	totalQuestions := 0

	if attempts, ok := doc["quizAttempt"].(map[string]interface{}); ok {
		for _, v := range attempts {
			attempt := v.(map[string]interface{})
			score := attempt["score"].(float64)
			totalScore += score
			if score == 4 {
				correctAnswers++
			} else if score == -1 {
				wrongAnswers++
			}

			if attempt["answer"].(string) != "" {
				totalQuestions++
			}
		}

	}

	accuracy := 0.0
	if totalQuestions > 0 {
		accuracy = float64(correctAnswers) / float64(totalQuestions)
	}

	doc["correctAnswers"] = correctAnswers
	doc["wrongAnswers"] = wrongAnswers
	doc["totalAttempted"] = totalQuestions
	doc["accuracy"] = accuracy
	doc["totalScore"] = totalScore

	fmt.Println(doc)

	update := bson.D{
		{
			Key:   "$set",
			Value: doc,
		},
	}

	options := options.Update().SetUpsert(false)

	result, err := quizBoardCollection.UpdateOne(ctx, bson.M{"_id": objBoard}, update, options)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CheckQuizBoardStatus(ctx context.Context, userId string, quizId string) (*models.QuizBoard, error) {
	result, err := GetQuizBoard(ctx, userId, quizId)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("New Document...")
		return nil, nil
	}

	token := result.QuizToken

	validToken, tokenErr := middlewares.ValidateToken(token)

	if validToken && result.IsActive && tokenErr != nil {
		return result, tokenErr
	} else if !validToken && result.IsActive && tokenErr != nil {
		_, err := UpdateQuizBoard(ctx, result.ID.Hex(), map[string]interface{}{
			"isActive": false,
		})
		if err != nil {
			return result, err
		}

		return result, tokenErr
	} else {
		return result, tokenErr
	}
}

func CreateQuizBoard(ctx context.Context, userId string, quizId string) (string, string, error) {
	_, err := CheckQuizBoardStatus(ctx, userId, quizId)

	if err != nil {
		return "", "", err
	}

	quiz, err := GetSingleQuiz(ctx, quizId)
	if err != nil {
		return "", "", nil
	}

	objQuiz, err := primitive.ObjectIDFromHex(quizId)
	if err != nil {
		return "", "", err
	}

	objUser, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return "", "", err
	}

	token, err := middlewares.CreateQuizToken(userId, quizId, quiz.Duration)
	if err != nil {
		return "", "", err
	}

	insertResult, insertErr := quizBoardCollection.InsertOne(ctx, models.QuizBoard{
		ID:        primitive.NewObjectID(),
		QuizId:    objQuiz,
		UserId:    objUser,
		QuizToken: token,
		IsActive:  true,
	})

	if insertErr != nil {
		return "", "", nil
	}

	fmt.Println(insertResult.InsertedID)

	return token, insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func GetResultBoard(ctx context.Context, boardId string) (*models.QuizBoard, error) {
	objBoard, _ := primitive.ObjectIDFromHex(boardId)

	var boardDetails *models.QuizBoard

	err := quizBoardCollection.FindOne(ctx, bson.M{"_id": objBoard}).Decode(&boardDetails)

	if err != nil {
		return nil, err
	}

	quiz, err := GetSingleQuiz(ctx, boardDetails.QuizId.Hex())

	if err != nil {
		return nil, err
	}

	boardDetails.QuizModel = quiz

	return boardDetails, nil
}

func RandomStudentBoard(ctx context.Context) (*mongo.InsertManyResult, error) {
	boards, err := utils.CreateRandomStudentForLeaderBoard()

	if err != nil {
		return nil, err
	}

	result, err := quizBoardCollection.InsertMany(ctx, boards)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateLeaderBoard(ctx context.Context, quizId string) ([]models.QuizBoard, error) {
	_quiz, err := primitive.ObjectIDFromHex(quizId)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"quizId": _quiz, "isOver": true}
	opts := options.Find().SetProjection(bson.M{"quizAttempt": 0, "quizToken": 0, "quizModel": 0}).SetSort(bson.D{{Key: "totalScore", Value: -1}, {Key: "accuracy", Value: -1}})

	var boards []models.QuizBoard

	cursor, err := quizBoardCollection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &boards); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	userChannel := make(chan *models.Users)

	for _, board := range boards {
		wg.Add(1)

		go func(userId string) {
			defer wg.Done()

			var user models.Users

			_user, err := primitive.ObjectIDFromHex(userId)

			if err != nil {
				return
			}

			err = userCollection.FindOne(ctx, bson.M{"_id": _user}).Decode(&user)

			if err != nil {
				return
			}
			userChannel <- &user
		}(board.UserId.Hex())
	}

	go func() {
		wg.Wait()
		close(userChannel)
	}()

	boardMap := make(map[string]*models.QuizBoard)

	for i := range boards {
		boardMap[boards[i].UserId.Hex()] = &boards[i]
	}

	for user := range userChannel {
		if board, ok := boardMap[user.ID.Hex()]; ok {
			board.User = *user
		}
	}
	return boards, nil
}
