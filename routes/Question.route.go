package routes

import (
	"quizGo/controllers"

	"github.com/gofiber/fiber/v2"
)

func QuestionRouter(app fiber.Router) {
	questionRouter := app.Group("/question")

	questionRouter.Post("/", controllers.CreateQuestionModel)
	questionRouter.Get("/", controllers.GetAllQuestionModel)
	questionRouter.Get("/:id", controllers.GetSingleQuestionModel)
	questionRouter.Put("/:id", controllers.UpdateQuestionModel)
	questionRouter.Delete("/:id", controllers.DeleteQuestionModel)
	questionRouter.Get("/topic/:id", controllers.GetAllTopicQuestionModel)
	questionRouter.Post("/multiple", controllers.InsertMultipleQuestionModel)
}
