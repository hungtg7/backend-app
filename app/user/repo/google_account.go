package repo

import (
	"context"

	"github.com/hungtg7/backend-app/app/user/entity"
)

func (r *UserRepoRepo) GetGoogleAccountByID(ctx context.Context, id string) (*entity.GoogleAccount, error) {
	var acc *entity.GetGoogleAccountByID

	query := r.Db

	if err := query.Where("google_id = ?", id).First(&pet).Error; err != nil {
		return nil, err
	}

	return pet, nil
}
