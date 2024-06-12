package controllers

import (
	"github.com/Pratham-Mishra04/interact-admin-microservice/config"
	"github.com/Pratham-Mishra04/interact-admin-microservice/helpers"
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

		case "project":
			var projects []models.Project
			if err := paginatedDB.
				Preload("User").
				Where("is_flagged=?", true).
				Order("created_at DESC").
				Find(&projects).Error; err != nil {
				return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
			}

			return c.Status(200).JSON(fiber.Map{
				"status":   "success",
				"message":  "",
				"projects": projects,
			})

		case "user":
			var users []models.User
			if err := paginatedDB.
				Where("is_flagged=?", true).
				Order("created_at DESC").
				Find(&users).Error; err != nil {
				return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
			}

			return c.Status(200).JSON(fiber.Map{
				"status":  "success",
				"message": "",
				"users":   users,
			})

		case "event":
			var events []models.Event
			if err := paginatedDB.
				Preload("Organization").
				Where("is_flagged=?", true).
				Order("created_at DESC").
				Find(&events).Error; err != nil {
				return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
			}

			return c.Status(200).JSON(fiber.Map{
				"status":  "success",
				"message": "",
				"events":  events,
			})

		case "opening":
			var openings []models.Opening
			if err := paginatedDB.
				Preload("Organization").
				Preload("Project").
				Preload("User").
				Where("is_flagged=?", true).
				Order("created_at DESC").
				Find(&openings).Error; err != nil {
				return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
			}

			return c.Status(200).JSON(fiber.Map{
				"status":   "success",
				"message":  "",
				"openings": openings,
			})

		case "announcement":
			var announcements []models.Announcement
			if err := paginatedDB.
				Preload("Organization").
				Where("is_flagged=?", true).
				Order("created_at DESC").
				Find(&announcements).Error; err != nil {
				return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
			}

			return c.Status(200).JSON(fiber.Map{
				"status":        "success",
				"message":       "",
				"announcements": announcements,
			})

		case "poll":
			var polls []models.Poll
			if err := paginatedDB.
				Preload("Organization").
				Where("is_flagged=?", true).
				Order("created_at DESC").
				Find(&polls).Error; err != nil {
				return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
			}

			return c.Status(200).JSON(fiber.Map{
				"status":  "success",
				"message": "",
				"polls":   polls,
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

			err = helpers.SendMailReq(comment.User.Email, 71, &comment.User, "comment", comment)
			if err != nil {
				return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
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

			err = helpers.SendMailReq(post.User.Email, 71, &post.User, "post", post)
			if err != nil {
				return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
			}

		case "user":
			var user models.User
			if err := initializers.DB.First(&user, "id = ?", parsedItemID).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return &fiber.Error{Code: 400, Message: "No User of this ID found."}
				}
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			user.IsFlagged = false

			if err := initializers.DB.Save(&user).Error; err != nil {
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			err = helpers.SendMailReq(user.Email, 71, &user, "user", nil)
			if err != nil {
				return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
			}

		case "project":
			var project models.Project
			if err := initializers.DB.First(&project, "id = ?", parsedItemID).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return &fiber.Error{Code: 400, Message: "No Project of this ID found."}
				}
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			project.IsFlagged = false

			if err := initializers.DB.Save(&project).Error; err != nil {
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			err = helpers.SendMailReq(project.User.Email, 71, &project.User, "project", project)
			if err != nil {
				return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
			}

		case "event":
			var event models.Event
			if err := initializers.DB.First(&event, "id = ?", parsedItemID).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return &fiber.Error{Code: 400, Message: "No Event of this ID found."}
				}
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			event.IsFlagged = false

			if err := initializers.DB.Save(&event).Error; err != nil {
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			err = helpers.SendMailReq(event.Organization.User.Email, 71, &event.Organization.User, "event", event)
			if err != nil {
				return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
			}

		case "opening":
			var opening models.Opening
			if err := initializers.DB.First(&opening, "id = ?", parsedItemID).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return &fiber.Error{Code: 400, Message: "No Opening of this ID found."}
				}
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			opening.IsFlagged = false

			if err := initializers.DB.Save(&opening).Error; err != nil {
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			err = helpers.SendMailReq(opening.User.Email, 71, &opening.User, "opening", opening)
			if err != nil {
				return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
			}

		case "announcement":
			var announcement models.Announcement
			if err := initializers.DB.First(&announcement, "id = ?", parsedItemID).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return &fiber.Error{Code: 400, Message: "No Announcement of this ID found."}
				}
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			announcement.IsFlagged = false

			if err := initializers.DB.Save(&announcement).Error; err != nil {
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			err = helpers.SendMailReq(announcement.Organization.User.Email, 71, &announcement.Organization.User, "announcement", announcement)
			if err != nil {
				return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
			}

		case "poll":
			var poll models.Poll
			if err := initializers.DB.First(&poll, "id = ?", parsedItemID).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return &fiber.Error{Code: 400, Message: "No Poll of this ID found."}
				}
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			poll.IsFlagged = false

			if err := initializers.DB.Save(&poll).Error; err != nil {
				return &fiber.Error{Code: 400, Message: config.DATABASE_ERROR}
			}

			err = helpers.SendMailReq(poll.Organization.User.Email, 71, &poll.Organization.User, "poll", poll)
			if err != nil {
				return &fiber.Error{Code: 500, Message: config.SERVER_ERROR}
			}
		}

		//TODO removed from flag email

		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Flag Removed",
		})
	}
}
