package controllers

import (
	"context"
	"fmt"
	"time"

	"quizGo/functions"
	"quizGo/handlers"
	"quizGo/middlewares"
	"quizGo/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUserModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		return handlers.InvalidBodyRouter(c)
	}

	result, err := functions.CreateUser(ctx, user)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to create user", err)
	}

	return handlers.SuccessRouter(c, "User created successfully", map[string]interface{}{
		"accessToken": result,
	})
}

func LoginUserModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user *models.Users

	if err := c.BodyParser(&user); err != nil {
		return handlers.InvalidBodyRouter(c)
	}

	result, err := functions.LoginUser(ctx, user.Email, user.Password)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to login user details", err)
	}

	return handlers.SuccessRouter(c, "Login Successful", map[string]interface{}{
		"accessToken": result,
	})
}

func GetUserModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := middlewares.GetUserId(c)

	if err != nil {
		return handlers.UserNotFoundRouter(c)
	}

	result, err := functions.GetUser(ctx, id)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to get user details", err)
	}

	return handlers.SuccessRouter(c, "User details found", &result)
}

func UpdateUserModel(c *fiber.Ctx) error {

	id, err := middlewares.GetUserId(c)

	if err != nil {
		return handlers.UserNotFoundRouter(c)
	}

	var user map[string]interface{}

	if err := c.BodyParser(&user); err != nil {
		return handlers.InvalidBodyRouter(c)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := functions.UpdateUser(ctx, id, user)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to update user details", err)
	}

	return handlers.SuccessRouter(c, "User updated successfully", result)
}

func DeleteUserModel(c *fiber.Ctx) error {
	id, err := middlewares.GetUserId(c)
	fmt.Println(id)

	if err != nil {
		return handlers.UserNotFoundRouter(c)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := functions.DeleteUser(ctx, id)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to delete user", err)
	}

	return handlers.SuccessRouter(c, "User deleted successfully", result)
}

func InsertMultipleUserModel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := functions.InsertMultipleUsers(ctx)

	if err != nil {
		return handlers.ErrorRouter(c, "Unable to create user", err)
	}

	return handlers.SuccessRouter(c, "User created successfully", result)
}
