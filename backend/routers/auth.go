package routers

import (
	"github.com/Pratham-Mishra04/interact-admin-microservice/controllers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app *fiber.App) {
	app.Post("/signup", middlewares.LimitRoute, controllers.SignUp)
	app.Post("/login", controllers.LogIn)
	app.Post("/refresh", controllers.Refresh)
}
