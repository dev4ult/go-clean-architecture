package repository

import (
	"gorm.io/gorm"

	entity "clean_arch/features/user/entity"
)

type Repository interface {
	Insert(input entity.UserEntity) (*entity.UserEntity, error)
	SelectAll() ([]entity.UserEntity, error)
	SelectById(userId int) (*entity.UserEntity, error)
	SelectByUsername(username string) (*entity.UserEntity, error)
}

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &userRepo {
		db: db,
	}
}