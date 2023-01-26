package entity

import (
	"time"
)

// ProductCategory entity.
type Order struct {
	ID      int32  `gorm:"primaryKey;auto_increment;column:id"`
	PetID   int32  `gorm:"foreignKey:column:pet_id"`
	Status  string `gorm:"column:status"`
	Address string `gorm:"column:address"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

// TableName returns ProductCategory's table name.
func (Order) TableName() string {
	return "order"
}
