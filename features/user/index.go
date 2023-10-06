package users

import (
	config "clean_arch/config"
	handler "clean_arch/features/user/handler"
	repo "clean_arch/features/user/repository"
	service "clean_arch/features/user/service"
)

func UserRouteHandler() handler.Handler {
	db := config.InitDB() 
	repo := repo.New(db)
	service := service.New(repo)
	return handler.New(service)
}