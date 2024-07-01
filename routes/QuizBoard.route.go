package routes

import (
	"quizGo/controllers"
	"quizGo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func QuizBoardRouter(c fiber.Router) {
	quizBoardRouter := c.Group("/board")

	quizBoardRouter.Use(middlewares.SetupJWT())
	quizBoardRouter.Post("/", controllers.CreateQuizBoardModel)
	quizBoardRouter.Get("/result/:id", controllers.GetResultBoardModel)
	quizBoardRouter.Post("/insert", controllers.RandomStudentBoardModel)
	quizBoardRouter.Get("/leaderboard/:id", controllers.GetLeaderBoardModel)
	quizBoardRouter.Get("/details/:id", controllers.UserStatusBoardModel)
	quizBoardRouter.Patch("/:id", controllers.UpdateBoardModel)

}
