package routers

import (
	"github.com/Pratham-Mishra04/interact-admin-microservice/controllers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/middlewares"
	"github.com/Pratham-Mishra04/interact-admin-microservice/models"
	"github.com/gofiber/fiber/v2"
)

func LogRouter(app *fiber.App) {
	app.Post("/logger/api", controllers.AddLog)

	logRoutes := app.Group("/logger", middlewares.Protect)
	logRoutes.Get("/", controllers.GetLogs)
	logRoutes.Post("/", controllers.AddLog)
	logRoutes.Get("/filter_data", controllers.GetFilterData)
	logRoutes.Delete("/:logID", middlewares.UserAuthorization(models.Manager), controllers.DeleteLog)
}
