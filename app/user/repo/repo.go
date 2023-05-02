package repo

import (
	"gorm.io/gorm"
)

type UserRepo struct{ Db *gorm.DB }

func New(db *gorm.DB) *UserRepo { return &UserRepo{db} }
