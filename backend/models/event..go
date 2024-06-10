package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Event struct {
	ID                  uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Title               string                `gorm:"type:text;not null" json:"title"`
	Tagline             string                `gorm:"type:text" json:"tagline"`
	CoverPic            string                `gorm:"type:text; default:default.jpg" json:"coverPic"`
	BlurHash            string                `gorm:"type:text; default:no-hash" json:"blurHash"`
	Description         string                `gorm:"type:text;not null" json:"description"`
	Links               pq.StringArray        `gorm:"type:text[]" json:"links"`
	Tags                pq.StringArray        `gorm:"type:text[]" json:"tags"`
	NoViews             int                   `gorm:"default:0" json:"noViews"`
	NoLikes             int                   `gorm:"default:0" json:"noLikes"`
	NoShares            int                   `gorm:"default:0" json:"noShares"`
	NoComments          int                   `gorm:"default:0" json:"noComments"`
	StartTime           time.Time             `gorm:"not null" json:"startTime"`
	EndTime             time.Time             `gorm:"not null" json:"endTime"`
	Location            string                `gorm:"not null" json:"location"`
	Category            string                `gorm:"type:text;not null" json:"category"`
	Impressions         int                   `gorm:"default:0" json:"noImpressions"`
	OrganizationID      uuid.UUID             `gorm:"type:uuid;not null" json:"organizationID"`
	Organization        Organization          `gorm:"" json:"organization"`
	CreatedAt           time.Time             `gorm:"default:current_timestamp" json:"createdAt"`
	IsFlagged   bool           `gorm:"default:false" json:"-"`
}
