package entity

import (
	"time"
)

// Produc tCategory entity.
type UserAccount struct {
	UserProfileID      int32  `gorm:"primaryKey:foreignKey:column:user_profile_id"`
	AccessToken        string `gorm:"foreignKey:column:access_token"`
	RefreshAccessToken string `gorm:"foreignKey:column:refresh_access_token"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

// TableName returns ProductCategory's table name.
func (UserAccount) TableName() string {
	return "user_account"
}
