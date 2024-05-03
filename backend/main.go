package main

import (
	"log"

	"github.com/Pratham-Mishra04/interact-admin-microservice/config"
	"github.com/Pratham-Mishra04/interact-admin-microservice/initializers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.ConnectToCache()
	initializers.AutoMigrate()

	config.AddLogger()
	// scripts.PopulateLogs()
}

func main() {
	defer config.LoggerCleanUp()
	app := fiber.New(fiber.Config{
		ErrorHandler: fiber.DefaultErrorHandler,
	})

	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(config.CORS())

	routers.Config(app)

	err := app.Listen(":" + initializers.CONFIG.PORT)
	if err != nil {
		log.Fatal(err)
	}
}
