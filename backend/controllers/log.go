package controllers

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/Pratham-Mishra04/interact-admin-microservice/config"
	"github.com/Pratham-Mishra04/interact-admin-microservice/initializers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/models"
	"github.com/Pratham-Mishra04/interact-admin-microservice/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddLog(c *fiber.Ctx) error {
	var reqBody models.LogEntrySchema
	if err := c.BodyParser(&reqBody); err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}

	go func(body models.LogEntrySchema) {
		var log models.Log

		log.Level = body.Level
		log.Title = body.Title
		log.Description = body.Description
		log.Path = body.Path
		log.Resource = models.RESOURCE(c.Get("Resource", ""))

		timestamp, err := time.Parse(time.RFC3339, body.Timestamp)
		if err == nil {
			log.Timestamp = timestamp
		}

		result := initializers.DB.Create(&log)
		if result.Error != nil {
			config.Logger.Errorw("Error while adding a log", "Error:", result.Error)
		} else {
			config.FlushCache()
		}
	}(reqBody)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Log Added",
	})
}

func GetLogs(c *fiber.Ctx) error {
	searchHash := getHashFromSearches(c)

	logsInCache := config.GetFromCache(searchHash)
	if logsInCache != nil {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Logs fetched",
			"logs":    logsInCache,
		})
	}

	paginatedDB := utils.Paginator(c)(initializers.DB)
	searchedDB := utils.Search(c)(paginatedDB)

	var logs []models.Log
	if err := searchedDB.
		Order("timestamp DESC").
		Find(&logs).Error; err != nil {
		return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
	}

	go config.SetToCache(searchHash, logs)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Logs fetched",
		"logs":    logs,
	})
}

func GetFilterData(c *fiber.Ctx) error {
	filterDataInCache := config.GetFilterDataFromCache("filterData")
	if filterDataInCache != nil {
		return c.Status(200).JSON(fiber.Map{
			"status":     "success",
			"message":    "Filter Data fetched",
			"filterData": filterDataInCache,
		})
	}

	filterData := models.FilterData{}

	fields := []string{"level", "path"}
	for _, field := range fields {
		if values, err := getAllUniqueValues(initializers.DB, field); err == nil {
			switch field {
			case "level":
				filterData.Levels = values
			case "path":
				filterData.Paths = values
			}
		}
	}

	go config.SetFilterDataToCache("filterData", filterData)

	return c.JSON(fiber.Map{
		"status":     "success",
		"message":    "",
		"filterData": filterData,
	})
}

func DeleteLog(c *fiber.Ctx) error {
	logID := c.Params("logID")

	var log models.Log
	if err := initializers.DB.First(&log, "id = ?", logID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &fiber.Error{Code: 400, Message: "No Log of this ID found."}
		}
		go config.Logger.Warn("Error while fetching log: ", "Error", err)
		return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
	}

	if err := initializers.DB.Delete(&log).Error; err != nil {
		go config.Logger.Warn("Error while deleting log: ", "Error", err)
		return &fiber.Error{Code: 500, Message: config.DATABASE_ERROR}
	}

	return c.Status(204).JSON(fiber.Map{
		"status":  "success",
		"message": "Log deleted successfully",
	})
}

func getAllUniqueValues(db *gorm.DB, field string) ([]string, error) {
	var values []string
	result := db.Model(&models.Log{}).Select(field).Group(field).Find(&values)
	if result.Error != nil {
		return nil, result.Error
	}
	return values, nil
}

func getHashFromSearches(c *fiber.Ctx) string {
	fields := []string{"title", "description", "path", "level", "start", "end", "limit", "page"}
	var values []string

	for _, field := range fields {
		values = append(values, c.Query(field, ""))
	}

	combinedString := strings.Join(values, ",")

	hash := sha256.New()
	hash.Write([]byte(combinedString))
	hashValue := fmt.Sprintf("%x", hash.Sum(nil))

	return hashValue
}
