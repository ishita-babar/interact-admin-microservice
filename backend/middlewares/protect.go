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

func verifyUserToken(tokenString string, user *models.LogUser) error {
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

func verifyAPIToken(tokenString string, SECRET string, resource models.RESOURCE) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		return &fiber.Error{Code: 403, Message: err.Error()}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return &fiber.Error{Code: 403, Message: "API Token is expired."}
		}
		jwtResource, ok := claims["sub"]
		if !ok || resource != jwtResource {
			return &fiber.Error{Code: 401, Message: "Invalid Resource in JWT claims."}
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
	err := verifyUserToken(tokenString, &user)
	if err != nil {
		return err
	}

	c.Set("loggedInUserID", fmt.Sprint(user.ID))
	c.Set("Resource", string(models.ADMIN))

	return c.Next()
}

func APIProtect(c *fiber.Ctx) error {
	authHeader := strings.Split(c.Get("Authorization"), " ")

	if len(authHeader) != 2 {
		return &fiber.Error{Code: 401, Message: "Not Authorized to use this API."}
	}

	jwtString := authHeader[1]

	apiToken := c.Get("API-TOKEN")
	if apiToken == "" {
		return &fiber.Error{Code: 403, Message: "Not Authorized to use this API."}
	}

	var err error

	switch apiToken {
	case initializers.CONFIG.BACKEND_TOKEN:
		err = verifyAPIToken(jwtString, initializers.CONFIG.BACKEND_SECRET, models.BACKEND)
		c.Set("Resource", string(models.BACKEND))

	case initializers.CONFIG.ML_TOKEN:
		err = verifyAPIToken(jwtString, initializers.CONFIG.ML_SECRET, models.ML)
		c.Set("Resource", string(models.ML))

	case initializers.CONFIG.SOCKETS_TOKEN:
		err = verifyAPIToken(jwtString, initializers.CONFIG.SOCKETS_SECRET, models.SOCKETS)
		c.Set("Resource", string(models.SOCKETS))

	case initializers.CONFIG.MAILER_TOKEN:
		err = verifyAPIToken(jwtString, initializers.CONFIG.MAILER_SECRET, models.MAILER)
		c.Set("Resource", string(models.MAILER))

	default:
		err = fmt.Errorf("not authorized to use this api, invalid api token")
	}

	if err != nil {
		return err
	}

	return c.Next()
}
