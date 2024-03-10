package models

import (
	"time"
)

type RESOURCE string

const (
	BACKEND RESOURCE = "backend"
	ML      RESOURCE = "ml"
	SOCKETS RESOURCE = "ws"
	MAILER  RESOURCE = "mailer"
	ADMIN   RESOURCE = "admin"
)

type Log struct {
	ID          int       `gorm:"autoIncrement;primaryKey" json:"id"`
	Level       string    `json:"level" gorm:"index:idx_level"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Path        string    `json:"path"`
	Resource    RESOURCE  `gorm:"index:idx_resource" json:"resource"`
	Timestamp   time.Time `json:"timestamp" gorm:"index:idx_timestamp"`
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
