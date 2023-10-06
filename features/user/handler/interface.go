package handler

import (
	"github.com/labstack/echo/v4"

	service "clean_arch/features/user/service"
)

type Handler interface {
	GetAllUsers() echo.HandlerFunc
	CreateUser() echo.HandlerFunc
}

type userController struct {
	service service.Service
}

func New (service service.Service) Handler {
	return &userController{
		service: service,
	}
}