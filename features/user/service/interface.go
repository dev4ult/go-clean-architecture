package service

import (
	dto "clean_arch/features/user/dtos"
	repo "clean_arch/features/user/repository"
)

type Service interface {
	FetchAll() ([]dto.UserResponse, string)
	CreateUser(input dto.UserInput) (*dto.UserResponse, string)
}

type userService struct {
	repo repo.Repository
}

func New(repo repo.Repository) Service {
	return &userService{
		repo: repo,
	}
}
