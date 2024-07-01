package routes

import "github.com/gofiber/fiber/v2"

func ServerRouter(app *fiber.App) {
	apiRouter := app.Group("/api")
	v1Router := apiRouter.Group("/v1")

	UserRoutes(v1Router)
	QuizRoutes(v1Router)
	QuestionRouter(v1Router)
	TopicRouter(v1Router)
	QuizBoardRouter(v1Router)
	SolutionRouter(v1Router)
}
