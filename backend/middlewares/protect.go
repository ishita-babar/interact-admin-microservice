package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/Pratham-Mishra04/interact-admin-microservice/config"
	"github.com/Pratham-Mishra04/interact-admin-microservice/initializers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func verifyToken(tokenString string, user *models.LogUser) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(initializers.CONFIG.JWT_SECRET), nil
	})

	if err != nil {
		return &fiber.Error{Code: 403, Message: config.TOKEN_EXPIRED_ERROR}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return &fiber.Error{Code: 403, Message: "Your token has expired."}
		}

		userID, ok := claims["sub"]
		if !ok {
			return &fiber.Error{Code: 401, Message: "Invalid user ID in token claims."}
		}

		if err := initializers.DB.First(user, "id = ?", userID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return &fiber.Error{Code: 401, Message: "User of this token no longer exists"}
			}
			return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
		}

		return nil
	} else {
		return &fiber.Error{Code: 403, Message: "Invalid Token"}
	}
}

func Protect(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenArr := strings.Split(authHeader, " ")

	if len(tokenArr) != 2 {
		return &fiber.Error{Code: 401, Message: "You are Not Logged In."}
	}

	tokenString := tokenArr[1]

	var user models.LogUser
	err := verifyToken(tokenString, &user)
	if err != nil {
		return err
	}

	return c.Next()
}
