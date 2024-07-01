package routes

import (
	"github.com/gofiber/fiber/v2"

	"quizGo/controllers"
	"quizGo/middlewares"
)

func UserRoutes(app fiber.Router) {
	userRouter := app.Group("/user")

	userRouter.Post("/", controllers.CreateUserModel)
	userRouter.Post("/login", controllers.LoginUserModel)
	userRouter.Use(middlewares.SetupJWT())
	userRouter.Get("/", controllers.GetUserModel)
	userRouter.Patch("/", controllers.UpdateUserModel)
	userRouter.Delete("/", controllers.DeleteUserModel)
	userRouter.Get("/insert",controllers.InsertMultipleUserModel)
}
