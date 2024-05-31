package controllers

import (
	"github.com/Pratham-Mishra04/interact-admin-microservice/config"
	"github.com/Pratham-Mishra04/interact-admin-microservice/initializers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/models"
	"github.com/Pratham-Mishra04/interact-admin-microservice/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetFlaggedItems(itemType string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		paginatedDB := utils.Paginator(c)(initializers.DB)

		switch itemType {
		case "comment":
			var comments []models.Comment
			if err := paginatedDB.
				Preload("User").
				Where("is_flagged=?", true).
				Order("created_at DESC").
				Find(&comments).Error; err != nil {
				return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
			}

			return c.Status(200).JSON(fiber.Map{
				"status":   "success",
				"message":  "",
				"comments": comments,
			})
		case "post":
			var posts []models.Post
			if err := paginatedDB.
				Preload("RePost").
				Preload("RePost.User").
				Preload("RePost.TaggedUsers").
				Preload("User").
				Preload("TaggedUsers").
				Where("is_flagged=?", true).
				Order("created_at DESC").
				Find(&posts).Error; err != nil {
				return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
			}

			return c.Status(200).JSON(fiber.Map{
				"status":  "success",
				"message": "",
				"posts":   posts,
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"status":  "failed",
			"message": config.SERVER_ERROR,
		})
	}
}

func RemoveFlag(itemType string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		parsedItemID, err := uuid.Parse(c.Params("itemID"))
		if err != nil {
			return &fiber.Error{Code: 400, Message: "Invalid ID"}
		}

		switch itemType {
		case "comment":
			var comment models.Comment
			if err := initializers.DB.First(&comment, "id = ?", parsedItemID).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return &fiber.Error{Code: 400, Message: "No Comment of this ID found."}
				}
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			comment.IsFlagged = false

			if err := initializers.DB.Save(&comment).Error; err != nil {
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

		case "post":
			var post models.Post
			if err := initializers.DB.First(&post, "id = ?", parsedItemID).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return &fiber.Error{Code: 400, Message: "No Post of this ID found."}
				}
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			post.IsFlagged = false

			if err := initializers.DB.Save(&post).Error; err != nil {
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}
		}

		//TODO removed from flag email

		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Flag Removed",
		})
	}
}
