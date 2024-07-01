package routes

import (
	"quizGo/controllers"

	"github.com/gofiber/fiber/v2"
)

func TopicRouter(app fiber.Router) {
	topicRouter := app.Group("/topic")

	topicRouter.Post("/", controllers.CreateTopicModel)
	topicRouter.Get("/", controllers.GetAllTopicModel)
	topicRouter.Put("/:id", controllers.UpdateTopicModel)
	topicRouter.Delete("/:id", controllers.DeleteTopicModel)
}
