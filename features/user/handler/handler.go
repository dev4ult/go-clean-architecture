package handler

import (
	"github.com/labstack/echo/v4"

	dto "clean_arch/features/user/dtos"
	helper "clean_arch/helpers"
)

func (uc *userController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {

		users, message := uc.service.FetchAll()

		if users == nil {
			return c.JSON(500, helper.Response(message, users))
		}

		return c.JSON(200, helper.Response(message, users))
	}
}

func (uc *userController) CreateUser() echo.HandlerFunc {
	var userInput dto.UserInput

	return func(c echo.Context) error {
		if err := c.Bind(&userInput); err != nil {
			return c.JSON(400, helper.Response(err.Error(), nil))
		}

		user, message := uc.service.CreateUser(userInput)
		
		if user == nil {
			return c.JSON(500, helper.Response(message, nil))
		}

		return c.JSON(200, helper.Response(message, user))
	}
}