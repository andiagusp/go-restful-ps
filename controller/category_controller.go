package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	FindAll(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindById(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
	Create(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(wr http.ResponseWriter, req *http.Request, params httprouter.Params)
}
