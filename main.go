package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	user_service "github.com/ninoloid/go-simple-rest/application/users/service"
	"github.com/ninoloid/go-simple-rest/common/helper"
	"github.com/ninoloid/go-simple-rest/infrastructure/database"
	user_controller "github.com/ninoloid/go-simple-rest/infrastructure/users/controller"
	user_repository "github.com/ninoloid/go-simple-rest/infrastructure/users/repository"
	users_router "github.com/ninoloid/go-simple-rest/infrastructure/users/router"
	"github.com/ninoloid/go-simple-rest/middleware"
)

func main() {

	database.ConnectDatabase()
	db := database.DB
	validate := validator.New()
	usersRepository := user_repository.NewUsersRepository(db)
	usersService := user_service.NewUsersService(usersRepository, validate)
	usersController := user_controller.NewUsersController(usersService)
	router := users_router.NewUserRouter(usersController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
