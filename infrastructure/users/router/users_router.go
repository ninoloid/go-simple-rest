package users_router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ninoloid/go-simple-rest/exception"
	user_controller "github.com/ninoloid/go-simple-rest/infrastructure/users/controller"
)

func NewUserRouter(usersController user_controller.UsersController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", usersController.FindAll)
	router.GET("/:userId", usersController.FindById)

	router.PanicHandler = exception.ErrorHandler

	return router
}
