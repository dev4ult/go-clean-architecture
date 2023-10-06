package main

import (
	"github.com/labstack/echo/v4"

	user "clean_arch/features/user"
	route "clean_arch/routes"
)

var (
	userHandler = user.UserRouteHandler()
)

func main() {
	e := echo.New()

	route.UserRoutes(e, userHandler)

	e.Start(":8000")
}