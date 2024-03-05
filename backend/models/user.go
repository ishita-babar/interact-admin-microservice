package models

import (
	"time"
)

type UserRole string

const (
	Member  UserRole = "Member"
	Manager UserRole = "Manager"
)

type LogUser struct {
	ID        int       `gorm:"autoIncrement;primaryKey" json:"id"`
	Username  string    `gorm:"type:text;unique;not null" json:"username"`
	Role      UserRole  `gorm:"type:text" json:"role"`
	Password  string    `json:"-"`
	CreatedAt time.Time `gorm:"default:current_timestamp;index:idx_created_at,sort:desc" json:"-"`
}

type UserCreateSchema struct {
	Username        string   `json:"username"`
	Role            UserRole `json:"role"`
	Password        string   `json:"password" validate:"required,min=8"`
	ConfirmPassword string   `json:"confirmPassword" validate:"required,min=8"`
}
