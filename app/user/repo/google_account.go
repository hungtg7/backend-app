package repo

import (
	"context"

	"github.com/hungtg7/backend-app/app/user/entity"
)

func (r *UserRepo) GetGoogleAccountByID(ctx context.Context, id string) (*entity.GoogleAccount, error) {
	var acc *entity.GoogleAccount

	query := r.Db

	if err := query.Where("google_id = ?", id).First(&acc).Error; err != nil {
		return nil, err
	}

	return acc, nil
}
