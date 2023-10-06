package service

import (
	dto "clean_arch/features/user/dtos"
	entity "clean_arch/features/user/entity"

	"github.com/mashingan/smapping"
)

func (us *userService) FetchAll() ([]dto.UserResponse, string) {
	users, err := us.repo.SelectAll()

	if err != nil {
		return nil, err.Error()
	}

	if len(users) < 1 {
		return []dto.UserResponse{}, "There is No User!"
	}

	var response []dto.UserResponse
	errMapping := smapping.FillStruct(&response, smapping.MapFields(&users))

	if errMapping != nil {
		return nil, "Something Went Wrong!"
	}

	// need to map entity to dto
	return response, "Success Get Users!"
}

func (us *userService) CreateUser(input dto.UserInput) (*dto.UserResponse, string) {
	user, err := us.repo.SelectByUsername(input.Username)

	if err != nil {
		return nil, err.Error()
	}

	if *user != (entity.User{}) {
		return nil, "User Has Already Exist!"
	}

	var userEntity entity.User
	errMappingDto := smapping.FillStruct(&userEntity, smapping.MapFields(&input))

	if errMappingDto != nil {
		panic(errMappingDto)
	}

	// need to map dto to entity
	newUser, errCreate := us.repo.Insert(userEntity)

	if errCreate != nil {
		return nil, "No User Created!"
	}

	var userDTO dto.UserResponse 
	errMappingEntity := smapping.FillStruct(&userDTO, smapping.MapFields(&newUser))

	if errMappingEntity != nil {
		panic(errMappingEntity)
	}

	// need to map entity to dto
	return &userDTO, "Success User Created!"
}