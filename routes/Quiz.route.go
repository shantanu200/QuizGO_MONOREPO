package routes

import (
	"quizGo/controllers"
	"quizGo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func QuizRoutes(app fiber.Router) {
	quizRouter := app.Group("/quiz")
	quizRouter.Use(middlewares.SetupJWT())
	quizRouter.Post("/", controllers.CreateQuizModel)
	quizRouter.Get("/", controllers.GetAllQuizzes)
	quizRouter.Put("/:id", controllers.UpdateQuizModel)
	quizRouter.Delete("/:id", controllers.DeleteQuizModel)
	quizRouter.Get("/:id", controllers.GetSingleQuizModel)
}
