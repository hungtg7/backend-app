package repo

import (
	"context"

	"github.com/hungtg7/backend-app/app/order/entity"
	"gorm.io/gorm"
)

type OrderRepo struct{ Db *gorm.DB }

func New(db *gorm.DB) *OrderRepo { return &OrderRepo{db} }

func (r *OrderRepo) GetOrderByID(ctx context.Context, id string) (*entity.Order, error) {
	var pet *entity.Order

	query := r.Db

	if err := query.Where("id = ?", id).First(&pet).Error; err != nil {
		return nil, err
	}

	return pet, nil
}

// Add adds new Pet to repo.
func (r *OrderRepo) Add(ctx context.Context, items ...*entity.Order) error {
	r.Db.CreateBatchSize = 100
	for _, item := range items {
		if err := r.Db.Create(item).Error; err != nil {
			return err
		}
	}

	r.Db.CreateBatchSize = 0
	return nil
}
