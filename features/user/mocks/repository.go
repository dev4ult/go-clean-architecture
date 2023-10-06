package mocks

import (
	entity "clean_arch/features/user/entity"
	"errors"
)

type MockRepo struct{}

func (mr *MockRepo) Insert(input entity.UserEntity) (*entity.UserEntity, error) {
	return &entity.UserEntity{Username: "Nibras"}, errors.New("No User Created!")
}

var ErrorSelect error

func (mr *MockRepo) SelectAll() ([]entity.UserEntity, error) {
	var users = []entity.UserEntity{}
	return users, ErrorSelect
}

func (mr *MockRepo) SelectById(userId int) (*entity.UserEntity, error) {
	var user entity.UserEntity

	return &user, nil
}


func (mr *MockRepo) SelectByUsername(username string) (*entity.UserEntity, error) {
	var users = []entity.UserEntity{
		{ID: 1, Username: "sarbinus", Email: "sarbin@example.com", Password: "sarbing123"},
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return &entity.UserEntity{}, nil
}

type SuccessMockRepo struct{}

func (mr *SuccessMockRepo) Insert(input entity.UserEntity) (*entity.UserEntity, error) {
	// if user, _ := mr.SelectByUsername("sarbinus"); user.Username == input.Username  {
	// 	return "User Has Already Exist", nil
	// } 
	return &input, nil
}

func (mr *SuccessMockRepo) SelectAll() ([]entity.UserEntity, error) {
	var users  = []entity.UserEntity{
		{ID: 1, Username: "sarbinus", Email: "sarbin@example.com", Password: "sarbin123"},
	}
	
	return users, nil
}

func (mr *SuccessMockRepo) SelectById(userId int) (*entity.UserEntity, error) {
	var user entity.UserEntity

	return &user, nil
}

func (mr *SuccessMockRepo) SelectByUsername(username string) (*entity.UserEntity, error) {
	var user = entity.UserEntity{}

	return &user, nil
}