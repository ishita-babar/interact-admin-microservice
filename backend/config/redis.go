package config

import (
	"context"
	"encoding/json"

	"github.com/Pratham-Mishra04/interact-admin-microservice/initializers"
	"github.com/Pratham-Mishra04/interact-admin-microservice/models"
	"github.com/redis/go-redis/v9"
)

var ctx = context.TODO()

func GetFromCache(key string) []models.Log {
	data, err := initializers.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			Logger.Warnw("Error Getting from cache", "Error:", err)
		}
		return nil
	}

	logs := []models.Log{}
	if err = json.Unmarshal([]byte(data), &logs); err != nil {
		Logger.Warnw("Error while unmarshaling logs", "Error:", err)
		return nil
	}
	return logs
}

func GetFilterDataFromCache(key string) *models.FilterData {
	data, err := initializers.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			Logger.Warnw("Error Getting from cache", "Error:", err)
		}
		return nil
	}

	filterData := models.FilterData{}
	if err = json.Unmarshal([]byte(data), &filterData); err != nil {
		Logger.Warnw("Error while unmarshaling filter data", "Error:", err)
		return nil
	}
	return &filterData
}

func SetToCache(key string, logs []models.Log) {
	data, err := json.Marshal(logs)
	if err != nil {
		Logger.Warnw("Error while marshaling logs", "Error:", err)
	}

	if err := initializers.RedisClient.Set(ctx, key, data, initializers.CacheExpirationTime).Err(); err != nil {
		Logger.Warnw("Error Setting to cache", "Error:", err)
	}
}

func SetFilterDataToCache(key string, filterData models.FilterData) {
	data, err := json.Marshal(filterData)
	if err != nil {
		Logger.Warnw("Error while marshaling filter data", "Error:", err)
	}

	if err := initializers.RedisClient.Set(ctx, key, data, initializers.CacheExpirationTime).Err(); err != nil {
		Logger.Warnw("Error Setting to filter data", "Error:", err)
	}
}

func RemoveFromCache(key string) {
	err := initializers.RedisClient.Del(ctx, key).Err()
	if err != nil && err != redis.Nil {
		Logger.Warnw("Error Removing from cache", "Error:", err)
	}
}

func FlushCache() {
	err := initializers.RedisClient.FlushAll(ctx).Err()
	if err != nil {
		Logger.Warnw("Error flushing cache", "Error", err)
	}
}
