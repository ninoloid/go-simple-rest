package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsersController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
