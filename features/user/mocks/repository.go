package mocks

import (
	entity "clean_arch/features/user/entity"
	"errors"
)

type MockRepo struct{}

func (mr *MockRepo) Insert(input entity.User) (*entity.User, error) {
	return &entity.User{Username: "Nibras"}, errors.New("No User Created!")
}

var ErrorSelect error

func (mr *MockRepo) SelectAll() ([]entity.User, error) {
	var users = []entity.User{}
	return users, ErrorSelect
}

func (mr *MockRepo) SelectById(userId int) (*entity.User, error) {
	var user entity.User

	return &user, nil
}


func (mr *MockRepo) SelectByUsername(username string) (*entity.User, error) {
	var users = []entity.User{
		{ID: 1, Username: "sarbinus", Email: "sarbin@example.com", Password: "sarbing123"},
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return &entity.User{}, nil
}

type SuccessMockRepo struct{}

func (mr *SuccessMockRepo) Insert(input entity.User) (*entity.User, error) {
	// if user, _ := mr.SelectByUsername("sarbinus"); user.Username == input.Username  {
	// 	return "User Has Already Exist", nil
	// } 
	return &input, nil
}

func (mr *SuccessMockRepo) SelectAll() ([]entity.User, error) {
	var users  = []entity.User{
		{ID: 1, Username: "sarbinus", Email: "sarbin@example.com", Password: "sarbin123"},
	}
	
	return users, nil
}

func (mr *SuccessMockRepo) SelectById(userId int) (*entity.User, error) {
	var user entity.User

	return &user, nil
}

func (mr *SuccessMockRepo) SelectByUsername(username string) (*entity.User, error) {
	var user = entity.User{}

	return &user, nil
}