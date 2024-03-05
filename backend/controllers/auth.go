package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/Pratham-Mishra04/interact-admin-microservice/config"
	"github.com/Pratham-Mishra04/interact-admin-microservice/initializers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateSendToken(c *fiber.Ctx, user models.LogUser, statusCode int, message string) error {
	access_token_claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"crt": time.Now().Unix(),
		"exp": time.Now().Add(config.ACCESS_TOKEN_TTL).Unix(),
	})

	refresh_token_claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"crt": time.Now().Unix(),
		"exp": time.Now().Add(config.REFRESH_TOKEN_TTL).Unix(),
	})

	access_token, err := access_token_claim.SignedString([]byte(initializers.CONFIG.JWT_SECRET))
	if err != nil {
		go config.Logger.Errorw("Error while decrypting JWT Token.", "Error:", err)
		return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
	}

	refresh_token, err := refresh_token_claim.SignedString([]byte(initializers.CONFIG.JWT_SECRET))
	if err != nil {
		go config.Logger.Errorw("Error while decrypting JWT Token.", "Error:", err)
		return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refresh_token,
		Expires:  time.Now().Add(config.REFRESH_TOKEN_TTL),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(statusCode).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"token":   access_token,
		"user":    user,
	})
}

func SignUp(c *fiber.Ctx) error {
	var reqBody models.UserCreateSchema

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Invalid Req Body"}
	}

	if reqBody.Password != reqBody.ConfirmPassword {
		return &fiber.Error{Code: 400, Message: "Passwords do not match."}
	}

	var user models.LogUser
	initializers.DB.First(&user, "username = ?", reqBody.Username)
	if user.ID != 0 {
		return &fiber.Error{Code: 400, Message: "User with this Username already exists"}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 12)
	if err != nil {
		go config.Logger.Errorw("Error while hashing Password.", "Error:", err)
		return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
	}

	newUser := models.LogUser{
		Password: string(hash),
		Username: reqBody.Username,
		Role:     reqBody.Role,
	}

	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		go config.Logger.Errorw("Error while adding a user", "Error:", result.Error)
		return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
	} else {
		return CreateSendToken(c, newUser, 201, "Account Created")
	}
}

func LogIn(c *fiber.Ctx) error {
	var reqBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Validation Failed"}
	}

	var user models.LogUser
	if err := initializers.DB.First(&user, "username = ? ", reqBody.Username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 400, Message: "No account with these credentials found."}
		} else {
			return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password)); err != nil {
		return &fiber.Error{Code: 400, Message: "No account with these credentials found."}
	}

	return CreateSendToken(c, user, 200, "Logged In")
}

func Refresh(c *fiber.Ctx) error {
	var reqBody struct {
		Token string `json:"token"`
	}

	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: "Validation Failed"}
	}

	access_token_string := reqBody.Token

	access_token, err := jwt.Parse(access_token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(initializers.CONFIG.JWT_SECRET), nil
	})

	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		go config.Logger.Infow("Token Expiration: ", "Error", err)
		return &fiber.Error{Code: 400, Message: config.TOKEN_EXPIRED_ERROR}
	}

	if access_token_claims, ok := access_token.Claims.(jwt.MapClaims); ok {

		access_token_userID, ok := access_token_claims["sub"]
		if !ok {
			return &fiber.Error{Code: 401, Message: "Invalid user ID in token claims."}
		}

		var user models.LogUser
		err := initializers.DB.First(&user, "id = ?", access_token_userID).Error
		if err != nil {
			go config.Logger.Warn("Error while fetching user for token refreshing", "Error:", err)
			return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
		}

		if user.ID == 0 {
			return &fiber.Error{Code: 401, Message: "User of this token no longer exists"}
		}

		refresh_token_string := c.Cookies("refresh_token")
		if refresh_token_string == "" {
			go config.Logger.Warn("Nil Refresh Token", "Error:", err)
			return &fiber.Error{Code: 401, Message: config.TOKEN_EXPIRED_ERROR}
		}

		refresh_token, err := jwt.Parse(refresh_token_string, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(initializers.CONFIG.JWT_SECRET), nil
		})

		if err != nil {
			go config.Logger.Infow("Token Expiration: ", "Error", err)
			return &fiber.Error{Code: 400, Message: config.TOKEN_EXPIRED_ERROR}
		}

		if refresh_token_claims, ok := refresh_token.Claims.(jwt.MapClaims); ok && refresh_token.Valid {
			refresh_token_userID, ok := refresh_token_claims["sub"]
			if !ok {
				return &fiber.Error{Code: 401, Message: "Invalid user ID in token claims."}
			}

			if refresh_token_userID != access_token_userID {
				go config.Logger.Warnw("Mismatched Tokens: ", "Access Token User ID", access_token_userID, "Refresh Token User ID", refresh_token_userID)
				return &fiber.Error{Code: 401, Message: "Mismatched Tokens."}
			}

			if time.Now().After(time.Unix(int64(refresh_token_claims["exp"].(float64)), 0)) {
				go config.Logger.Infow("Token Expiration: ", "Error", err)
				return &fiber.Error{Code: 401, Message: config.TOKEN_EXPIRED_ERROR}
			}

			new_access_token_claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": user.ID,
				"crt": time.Now().Unix(),
				"exp": time.Now().Add(config.ACCESS_TOKEN_TTL).Unix(),
			})

			new_access_token, err := new_access_token_claim.SignedString([]byte(initializers.CONFIG.JWT_SECRET))
			if err != nil {
				go config.Logger.Warn("Error while decrypting JWT Token.: ", "Error", err)
				return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
			}

			return c.Status(200).JSON(fiber.Map{
				"status": "success",
				"token":  new_access_token,
			})
		}

		return nil
	} else {
		return &fiber.Error{Code: 401, Message: "Invalid Token"}
	}
}
