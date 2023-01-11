package entity

import (
	"time"
)

// ProductCategory entity.
type Pet struct {
	ID   string `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

// TableName returns ProductCategory's table name.
func (Pet) TableName() string {
	return "pet"
}
