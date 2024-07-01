package controllers

import (
	"context"
	"quizGo/functions"
	"quizGo/handlers"
	"quizGo/middlewares"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateQuizBoardModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_id, err := middlewares.GetUserId(c)
	if err != nil {
		return handlers.UserNotFoundRouter(c)
	}

	var doc map[string]interface{}

	if err := c.BodyParser(&doc); err != nil {
		return handlers.InvalidBodyRouter(c)
	}

	token, boardId, err := functions.CreateQuizBoard(ctx, _id, doc["quizId"].(string))
	if err != nil {
		return handlers.ErrorRouter(c, "Unable to attempt quiz", err)
	}

	return handlers.SuccessRouter(c, "Quiz connected to user", fiber.Map{
		"quizToken": token,
		"userId":    _id,
		"quizId":    doc["quizId"],
		"boardId":   boardId,
	})
}

func UserStatusBoardModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_userId, err := middlewares.GetUserId(c)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to access userId", err)
	}

	quizId := c.Params("id")

	result, err := functions.CheckQuizBoardStatus(ctx, _userId, quizId)

	if err == nil {
		return handlers.SuccessRouter(c, "New Quiz", map[string]interface{}{"isActive": false, "isStart": false, "isOver": false})
	}

	return handlers.SuccessRouter(c, "Resume Quiz", map[string]interface{}{"isStart": true, "isActive": result.IsActive, "boardId": result.ID, "isOver": result.IsOver, "totalScore": result.TotalScore})
}

func UpdateBoardModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var updateBody map[string]interface{}

	if err := c.BodyParser(&updateBody); err != nil {
		return handlers.InvalidBodyRouter(c)
	}

	boardId := c.Params("id")

	result, err := functions.UpdateQuizBoard(ctx, boardId, updateBody)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to update quiz details", err)
	}

	return handlers.SuccessRouter(c, "Quiz details updated", result)
}

func GetResultBoardModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	result, err := functions.GetResultBoard(ctx, id)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to get quiz result for user", err)
	}

	return handlers.SuccessRouter(c, "Result found successfully", result)
}

func RandomStudentBoardModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := functions.RandomStudentBoard(ctx)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to process data", err)
	}

	return handlers.SuccessRouter(c, "Result found successfully", result)
}

func GetLeaderBoardModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	quizId := c.Params("id")
	result, err := functions.CreateLeaderBoard(ctx, quizId)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to get leaderboard details", err)
	}

	return handlers.SuccessRouter(c, "Leaderboard found", result)
}
