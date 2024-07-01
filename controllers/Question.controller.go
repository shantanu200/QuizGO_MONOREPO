package controllers

import (
	"context"
	"fmt"
	"quizGo/functions"
	"quizGo/handlers"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateQuestionModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var question *map[string]interface{}

	if err := c.BodyParser(&question); err != nil {
		fmt.Println(err.Error())
		return handlers.InvalidBodyRouter(c)
	}

	result, err := functions.CreateQuestion(ctx, question)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to create question", err)
	}

	return handlers.SuccessRouter(c, "Question created successfully", result)
}

func GetAllQuestionModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	query := c.Query("q")

	result, err := functions.GetQuestions(ctx, int64(page), int64(limit), query)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to get questions", err)
	}

	return handlers.SuccessRouter(c, "Question found successfully", result)
}

func GetSingleQuestionModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	result, err := functions.SingleQuestion(ctx, id)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to get question details", err)
	}

	return handlers.SuccessRouter(c, "Question details found", result)
}

func UpdateQuestionModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	var updateBody map[string]interface{}

	if err := c.BodyParser(&updateBody); err != nil {
		return handlers.InvalidBodyRouter(c)
	}

	result, err := functions.UpdateQuestion(ctx, id, updateBody)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to update question details", err)
	}

	return handlers.SuccessRouter(c, "Question details updated successfully", result)
}

func DeleteQuestionModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	result, err := functions.DeleteQuestion(ctx, id)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to delete question", err)
	}

	return handlers.SuccessRouter(c, "Question details deleted", result)
}

func GetAllTopicQuestionModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	query := c.Query("q")

	result, err := functions.GetTopicWiseQuestion(ctx, id, int64(page), int64(limit), query)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to get topic wise questions", err)
	}

	return handlers.SuccessRouter(c, "All topic wise questions", result)
}

func InsertMultipleQuestionModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var question []interface{}

	if err := c.BodyParser(&question); err != nil {
		fmt.Println(err.Error())
		return handlers.InvalidBodyRouter(c)
	}

	result, err := functions.InsertMultipleQuestion(ctx, question)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to create question", err)
	}

	return handlers.SuccessRouter(c, "Question created successfully", result)
}
