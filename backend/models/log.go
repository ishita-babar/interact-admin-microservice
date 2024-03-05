package models

import (
	"time"
)

type Log struct {
	ID          int    `gorm:"autoIncrement;primaryKey" json:"id"`
	Level       string `json:"level" gorm:"index:idx_level"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Path        string `json:"path"`
	// Resource       string    `json:"resourceId" gorm:"index:idx_resource_id"`
	Timestamp time.Time `json:"timestamp" gorm:"index:idx_timestamp"`
}

type LogEntrySchema struct {
	Level       string `json:"level"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Timestamp   string `json:"timestamp"`
}

type FilterData struct {
	Levels []string `json:"levels"`
	Paths  []string `json:"paths"`
}
