package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ninoloid/go-simple-rest/application/users/service"
	"github.com/ninoloid/go-simple-rest/common/helper"
	"github.com/ninoloid/go-simple-rest/infrastructure/database"
	"github.com/ninoloid/go-simple-rest/infrastructure/users/controller"
	"github.com/ninoloid/go-simple-rest/infrastructure/users/repository"
	users_router "github.com/ninoloid/go-simple-rest/infrastructure/users/router"
	"github.com/ninoloid/go-simple-rest/middleware"
)

func main() {

	database.ConnectDatabase()
	db := database.DB
	validate := validator.New()
	usersRepository := repository.NewUsersRepository(db)
	usersService := service.NewUsersService(usersRepository, validate)
	usersController := controller.NewUsersController(usersService)
	router := users_router.NewUserRouter(usersController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
