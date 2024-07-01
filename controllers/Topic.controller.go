package controllers

import (
	"context"
	"quizGo/functions"
	"quizGo/handlers"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateTopicModel(c *fiber.Ctx) error {
	var doc map[string]interface{}

	if err := c.BodyParser(&doc); err != nil {
		return handlers.InvalidBodyRouter(c)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := functions.CreateTopic(ctx, doc)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to create new topic", err)
	}

	return handlers.SuccessRouter(c, "New topic created", result)
}

func GetAllTopicModel(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	query := c.Query("q")

	result, err := functions.GetAllTopics(ctx, int64(page), int64(limit), query)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to get topics", err)
	}

	return handlers.SuccessRouter(c, "All topics details", result)
}

func UpdateTopicModel(c *fiber.Ctx) error {
	var doc map[string]interface{}

	if err := c.BodyParser(&doc); err != nil {
		return handlers.InvalidBodyRouter(c)
	}

	id := c.Params("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := functions.UpdateTopicDetails(ctx, id, doc)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to update topic details", err)
	}

	return handlers.SuccessRouter(c, "Topic details updated", result)
}

func DeleteTopicModel(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := functions.DeleteTopicDetails(ctx, id)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to delete topic details", err)
	}

	return handlers.SuccessRouter(c, "Topic details deleted", result)
}
