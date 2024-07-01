package main

import (
	"log"
	"quizGo/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(healthcheck.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 10 * time.Second,
	}))
	routes.ServerRouter(app)
	log.Fatal(app.Listen("127.0.0.1:9000"))
}
