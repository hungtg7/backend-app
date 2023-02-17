package repo

import (
	"context"

	"github.com/hungtg7/backend-app/app/pet/entity"
	"gorm.io/gorm"
)

type PetRepo struct{ Db *gorm.DB }

func New(db *gorm.DB) *PetRepo { return &PetRepo{db} }

func (r *PetRepo) GetPetByID(ctx context.Context, id string) (*entity.Pet, error) {
	var pet *entity.Pet

	query := r.Db

	if err := query.Where("id = ?", id).First(&pet).Error; err != nil {
		return nil, err
	}

	return pet, nil
}

func (r *PetRepo) CountAllPet(ctx context.Context) int64 {
	var count int64
	r.Db.Table("pet").Count(&count)
	return count
}

func (r *PetRepo) GetPets(ctx context.Context, offset int, limit int) ([]*entity.Pet, error) {
	var pet []*entity.Pet

	query := r.Db
	if err := query.Limit(limit).Offset(offset).Find(pet).Error; err != nil {
		return nil, err
	}

	return pet, nil
}

// Add adds new Pet to repo.
func (r *PetRepo) Add(ctx context.Context, items ...*entity.Pet) error {
	r.Db.CreateBatchSize = 100
	for _, item := range items {
		if err := r.Db.Create(item).Error; err != nil {
			return err
		}
	}

	r.Db.CreateBatchSize = 0
	return nil
}
