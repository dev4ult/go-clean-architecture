package repository

import (
	"gorm.io/gorm"

	entity "clean_arch/features/user/entity"
)

type Repository interface {
	Insert(input entity.User) (*entity.User, error)
	SelectAll() ([]entity.User, error)
	SelectById(userId int) (*entity.User, error)
	SelectByUsername(username string) (*entity.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &userRepo {
		db: db,
	}
}