package routes

import (
	"quizGo/controllers"
	"quizGo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SolutionRouter(c fiber.Router) {
	solutionRouter := c.Group("/solution")
	solutionRouter.Use(middlewares.SetupJWT())
	solutionRouter.Post("/", controllers.CreateSolutionModel)
	solutionRouter.Get("/:id", controllers.GetSolutionBoard)
}
