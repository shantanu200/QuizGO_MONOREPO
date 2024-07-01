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

func CreateQuizModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var quiz *map[string]interface{}

	if err := c.BodyParser(&quiz); err != nil {
		fmt.Println(err.Error())
		return handlers.InvalidBodyRouter(c)
	}

	result, err := functions.CreateQuiz(ctx, quiz)
	if err != nil {
		return handlers.ErrorRouter(c, "Unable to create quiz", err)
	}

	return handlers.SuccessRouter(c, "Quiz created successfully", result)
}

func GetAllQuizzes(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	query := c.Query("q")

	result, err := functions.GetAllQuizzes(ctx, int64(page), int64(limit), query)
	if err != nil {
		return handlers.ErrorRouter(c, "Unable to get quizzess", err)
	}

	return handlers.SuccessRouter(c, "All Quizzess", result)
}

func UpdateQuizModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var quiz map[string]interface{}

	if err := c.BodyParser(&quiz); err != nil {
		fmt.Println(err.Error())
		return handlers.InvalidBodyRouter(c)
	}

	id := c.Params("id")

	result, err := functions.UpdateQuiz(ctx, id, quiz)
	if err != nil {
		return handlers.ErrorRouter(c, "Unable to update quiz", err)
	}

	return handlers.SuccessRouter(c, "Quiz updated successfully", result)
}

func DeleteQuizModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	result, err := functions.DeleteQuiz(ctx, id)
	if err != nil {
		return handlers.ErrorRouter(c, "Unable to delete quiz", err)
	}

	return handlers.SuccessRouter(c, "Quiz deleted successfully", result)
}

func GetSingleQuizModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")

	result, err := functions.GetSingleQuiz(ctx, id)
	if err != nil {
		return handlers.ErrorRouter(c, "Unable to get quiz", err)
	}

	return handlers.SuccessRouter(c, "Quiz details found successfully", result)
}
