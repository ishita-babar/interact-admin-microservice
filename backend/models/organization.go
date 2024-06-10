package models

import (
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID                uuid.UUID                `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	UserID            uuid.UUID                `gorm:"type:uuid;not null" json:"userID"` //user model who is given the organization status
	User              User                     `gorm:"" json:"user"`
	OrganizationTitle string                   `gorm:"unique" json:"title"`
	Invitations       []Invitation             `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE" json:"invitations"`
	Events            []Event                  `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE" json:"events"`
	NumberOfMembers   int16                    `gorm:"default:0" json:"noMembers"`
	NumberOfEvents    int16                    `gorm:"default:0" json:"noEvents"`
	NumberOfProjects  int16                    `gorm:"default:0" json:"noProjects"`
	CreatedAt         time.Time                `gorm:"default:current_timestamp" json:"createdAt"`
	IsFlagged   bool           `gorm:"default:false" json:"-"`
}
