package mocks

import (
	dto "clean_arch/features/user/dtos"
	entity "clean_arch/features/user/entity"
)

type MockService struct{}

var UsersResponse = []dto.UserResponse{}

func (us *MockService) FetchAll() ([]dto.UserResponse, string) {
	if len(UsersResponse) < 1 {
		return UsersResponse, "There is No User!"
	}

	return nil, "Something Went Wrong!"
}

// binding


func (us *MockService) CreateUser(input dto.UserInput) (*dto.UserResponse, string) {
	users := []entity.UserEntity{
		{ID: 1, Username: "sarbinus", Email: "sarbin@example.com", Password: "sarbin123"},
	}

	for _, user := range users {
		if user.Username == input.Username {
			return nil, "User Has Already Exist!"
		}
	}

	return nil, "Something Went Wrong!"
}

type SuccessMockService struct{}

func (us *SuccessMockService) FetchAll() ([]dto.UserResponse, string) {
	return []dto.UserResponse{
		{Username: "sarbinus", Email: "sarbin@example.com"},
	}, "Success Get Users!"
}

func (us *SuccessMockService) CreateUser(input dto.UserInput) (*dto.UserResponse, string) {
	return &dto.UserResponse {
			Username: input.Username, 
			Email: input.Email,
		}, "Success User Created!"
}