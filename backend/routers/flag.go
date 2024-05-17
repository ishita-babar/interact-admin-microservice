package routers

import (
	"github.com/Pratham-Mishra04/interact-admin-microservice/controllers"
	"github.com/gofiber/fiber/v2"
)

func FlagRouter(app *fiber.App) {
	flagRoutes := app.Group("/flags")
	flagRoutes.Get("/comments", controllers.GetFlaggedItems("comment"))
	flagRoutes.Get("/posts", controllers.GetFlaggedItems("post"))

	flagRoutes.Post("/comments/:itemID", controllers.RemoveFlag("comment"))
	flagRoutes.Post("/posts/:itemID", controllers.RemoveFlag("post"))
}
