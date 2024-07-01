package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SuccessRouter(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func ErrorRouter(c *fiber.Ctx, message string, err error) error {
	fmt.Println("Error: ", err.Error())

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": message,
		"error":   string(err.Error()),
	})
}

func ServerRouter(c *fiber.Ctx, err error) error {
	fmt.Println("Server Error: ", err.Error())

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"success": false,
		"message": "Internal Server Error",
	})
}

func UserNotFoundRouter(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"success": false,
		"message": "User not found | Please login with valid details",
	})
}

func InvalidBodyRouter(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": "Invalid Body Passed",
	})
}
