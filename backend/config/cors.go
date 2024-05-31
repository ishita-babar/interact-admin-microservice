package config

import (
	"fmt"

	"github.com/Pratham-Mishra04/interact-admin-microservice/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("%s, %s, %s, %s, %s",
			initializers.CONFIG.FRONTEND_URL,
			initializers.CONFIG.BACKEND_URL,
			initializers.CONFIG.MAILER_URL,
			initializers.CONFIG.ML_URL,
			initializers.CONFIG.WS_URL),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	})
}
