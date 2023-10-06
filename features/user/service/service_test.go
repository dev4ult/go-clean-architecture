package service

import (
	dto "clean_arch/features/user/dtos"
	mock "clean_arch/features/user/mocks"
	"errors"
	"testing"
)

func TestFetchAll(t *testing.T) {
	var repo = mock.MockRepo{}
	var service = New(&repo)

	if _, message := service.FetchAll(); message != "There is No User!" {
		t.Error("Any user should not exist!")
	}

	mock.ErrorSelect = errors.New("Error SELECT FROM users!")
	if _, message := service.FetchAll(); message != "Error SELECT FROM users!" {
		t.Error("There is should be select error!", message)
	}

	var successRepo = mock.SuccessMockRepo{}
	var successService = New(&successRepo)

	if _, message := successService.FetchAll(); message != "Success Get Users!" {
		t.Error("User should already be listed!")
	}
}

func TestCreateUser(t *testing.T)  {
	var repo = mock.MockRepo{}
	var service = New(&repo)

	var existedUser = dto.UserInput{Username: "sarbinus", Email: "sarbin@example.com", Password: "sarbing123"}
	
	if _, message := service.CreateUser(existedUser); message != "User Has Already Exist!" {
		t.Error("User should aready exist!")
	}

	var emptyInput = dto.UserInput{}
	if user, _ := service.CreateUser(emptyInput); user != nil {
		t.Error("User should be nil!")
	}
	
	var successRepo = mock.SuccessMockRepo{}
	var successService = New(&successRepo)
	var nonExistedUser = dto.UserInput{Username: "nibras", Email: "nibras@example.com", Password: "nibras123"}

	if _, message := successService.CreateUser(nonExistedUser); message != "Success User Created!" {
		t.Error("User should be successfully created!")
	}
}