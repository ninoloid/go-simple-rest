package user_controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ninoloid/go-simple-rest/application/users/response"
	user_service "github.com/ninoloid/go-simple-rest/application/users/service"
	"github.com/ninoloid/go-simple-rest/common/helper"
)

type UsersControllerImpl struct {
	UsersService user_service.UsersService
}

func NewUsersController(userService user_service.UsersService) UsersController {
	return &UsersControllerImpl{
		UsersService: userService,
	}
}

func (controller *UsersControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UsersService.FindById(request.Context(), id)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsersControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := controller.UsersService.FindAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
