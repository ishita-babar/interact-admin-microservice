package middlewares

import (
	"github.com/Pratham-Mishra04/interact-admin-microservice/initializers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/models"
	"github.com/gofiber/fiber/v2"
)

func checkAccess(UserRole models.UserRole, AuthorizedRole models.UserRole) bool {
	if UserRole == models.Manager {
		return true
	} else if UserRole == models.Member {
		return AuthorizedRole == models.Member
	}

	return false
}

func UserAuthorization(Role models.UserRole) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		loggedInUserID := c.GetRespHeader("loggedInUserID")

		var user models.LogUser
		err := initializers.DB.First(&user, "id = ?", loggedInUserID).Error
		if err != nil {
			return &fiber.Error{Code: 400, Message: "No user of this id found."}
		}

		if !checkAccess(user.Role, Role) {
			return &fiber.Error{Code: 403, Message: "You don't have the Permission to perform this action."}
		}

		return c.Next()
	}
}

func LimitRoute(c *fiber.Ctx) error {
	return &fiber.Error{Code: 403, Message: "Not Accepting Users."}
}
