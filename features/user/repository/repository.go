package repository

import entity "clean_arch/features/user/entity"

func (ur *userRepo) Insert(input entity.UserEntity) (*entity.UserEntity, error) {
	if err := ur.db.Create(input).Error; err != nil {
		return nil, err
	}

	return &input, nil
}

func (ur *userRepo) SelectAll() ([]entity.UserEntity, error) {
	var users []entity.UserEntity
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepo) SelectById(userId int) (*entity.UserEntity, error) {
	var user entity.UserEntity

	if err := ur.db.First(&user, userId).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepo) SelectByUsername(username string) (*entity.UserEntity, error) {
	var user entity.UserEntity

	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}