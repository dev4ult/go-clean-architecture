package routes

import (
	user "clean_arch/features/user/handler"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, handler user.Handler) {
	users := e.Group("/users")

	users.GET("", handler.GetAllUsers())
	users.POST("", handler.CreateUser())
}