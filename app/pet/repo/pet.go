package repo

import (
	"context"
	"fmt"

	"github.com/hungtg7/api-app/app/pet/entity"
	"gorm.io/gorm"
)

type PetRepo struct{ db *gorm.DB }

func (r *PetRepo) GetPetByID(ctx context.Context, Id string) (*entity.Pet, error) {
	return nil, fmt.Errorf("error")
}
