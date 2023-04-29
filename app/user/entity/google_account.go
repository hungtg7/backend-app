package entity

import (
	"time"
)

type GoogleAccount struct {
	GoogleID      string `gorm:"primaryKey:column:google_id"`
	UserProfileID int32  `gorm:"foreignKey:column:user_profile_id"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

// TableName returns ProductCategory's table name.
func (UserProfile) TableName() string {
	return "google_account"
}
