package entity

import (
	"time"
)

// Produc tCategory entity.
type UserProfile struct {
	ID        int32  `gorm:"primaryKey;auto_increment;column:id"`
	FirstName string `gorm:"foreignKey:column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

// TableName returns ProductCategory's table name.
func (UserProfile) TableName() string {
	return "user_profile"
}
