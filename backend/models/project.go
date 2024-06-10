package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Project struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Title           string         `gorm:"type:text;not null" json:"title"` //TODO31 Validation error handling for no of chars
	Slug            string         `gorm:"type:text;not null" json:"slug"`
	Tagline         string         `gorm:"type:text;not null" json:"tagline"`
	CoverPic        string         `gorm:"type:text; default:default.jpg" json:"coverPic"`
	BlurHash        string         `gorm:"type:text; default:no-hash" json:"blurHash"`
	Description     string         `gorm:"type:text;not null" json:"description"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null" json:"userID"`
	User            User           `gorm:"" json:"user"`
	CreatedAt       time.Time      `gorm:"default:current_timestamp" json:"createdAt"`
	Tags            pq.StringArray `gorm:"type:text[]" json:"tags"`
	NoLikes         int            `gorm:"default:0" json:"noLikes"`
	NoShares        int            `gorm:"default:0" json:"noShares"`
	NoComments      int            `gorm:"default:0" json:"noComments"`
	TotalNoViews    int            `gorm:"default:0" json:"totalNoViews"`
	Category        string         `gorm:"type:text;not null" json:"category"`
	IsPrivate       bool           `gorm:"default:false" json:"isPrivate"`
	Views           int            `json:"views"`
	NumberOfMembers int            `gorm:"default:1" json:"noMembers"`
	Impressions     int            `gorm:"default:0" json:"noImpressions"`
	Links           pq.StringArray `gorm:"type:text[]" json:"links"`
	PrivateLinks    pq.StringArray `gorm:"type:text[]" json:"-"`
	IsFlagged   bool           `gorm:"default:false" json:"-"`
}