package controllers

import (
	"context"
	"quizGo/functions"
	"quizGo/handlers"
	"quizGo/middlewares"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateSolutionModel(c *fiber.Ctx) error {
	userId, err := middlewares.GetUserId(c)
	if err != nil {
		return handlers.ErrorRouter(c, "Invalid Middleware Request", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var solution map[string]interface{}

	if err := c.BodyParser(&solution); err != nil {
		return handlers.InvalidBodyRouter(c)
	}

	result, err := functions.CreateSolution(ctx, userId, solution)
	if err != nil {
		return handlers.ErrorRouter(c, "Unable to create solution", err)
	}

	return handlers.SuccessRouter(c, "Solution Created", result)
}

func GetSolutionBoard(c *fiber.Ctx) error {
	userId, err := middlewares.GetUserId(c)

	boardId := c.Params("id")

	if err != nil {
		return handlers.ErrorRouter(c, "Invalid Middleware Request", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := functions.GetSolution(ctx, userId, boardId)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to fetch user solution", err)
	}

	return handlers.SuccessRouter(c, "User Solution Fetched", result)
}
