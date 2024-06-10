package routers

import (
	"github.com/Pratham-Mishra04/interact-admin-microservice/controllers"
	"github.com/gofiber/fiber/v2"
)

func FlagRouter(app *fiber.App) {
	flagRoutes := app.Group("/flags")
	flagRoutes.Get("/comments", controllers.GetFlaggedItems("comment"))
	flagRoutes.Get("/posts", controllers.GetFlaggedItems("post"))
	flagRoutes.Get("/users", controllers.GetFlaggedItems("user"))
	flagRoutes.Get("/projects", controllers.GetFlaggedItems("project"))
	flagRoutes.Get("/events", controllers.GetFlaggedItems("event"))
	flagRoutes.Get("/openings", controllers.GetFlaggedItems("opening"))
	flagRoutes.Get("/announcements", controllers.GetFlaggedItems("announcement"))
	flagRoutes.Get("/polls", controllers.GetFlaggedItems("poll"))

	flagRoutes.Post("/comments/:itemID", controllers.RemoveFlag("comment"))
	flagRoutes.Post("/posts/:itemID", controllers.RemoveFlag("post"))
	flagRoutes.Post("/users/:itemID", controllers.RemoveFlag("user"))
	flagRoutes.Post("/projects/:itemID", controllers.RemoveFlag("project"))
	flagRoutes.Post("/events/:itemID", controllers.RemoveFlag("event"))
	flagRoutes.Post("/openings/:itemID", controllers.RemoveFlag("opening"))
	flagRoutes.Post("/announcements/:itemID", controllers.RemoveFlag("announcement"))
	flagRoutes.Post("/polls/:itemID", controllers.RemoveFlag("poll"))
}
