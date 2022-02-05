package controller

import (
	"golang-restful/helper"
	"golang-restful/model/web"
	"golang-restful/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}

func (controller *CategoryControllerImpl) FindAll(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	categories := controller.CategoryService.FindAll()

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Get All Category",
		Data:   categories,
	}

	helper.WriteEncodeJson(wr, webResponse)
}

func (controller *CategoryControllerImpl) FindById(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.Atoi(param)
	helper.PanicHandler(err)

	category := controller.CategoryService.FindById(id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Get Category",
		Data:   category,
	}

	helper.WriteEncodeJson(wr, webResponse)
}

func (controller *CategoryControllerImpl) Create(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var ccr web.CategoryCreateRequest
	helper.DecodeJson(req, &ccr)
	cr := controller.CategoryService.Create(ccr)

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Success Create Category",
		Data:   cr,
	}

	helper.WriteEncodeJson(wr, webResponse)
}

func (controller *CategoryControllerImpl) Update(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var cur web.CategoryUpdateRequest
	helper.DecodeJson(req, &cur)

	param := params.ByName("id")
	id, err := strconv.Atoi(param)
	helper.PanicHandler(err)

	cur.Id = id
	category := controller.CategoryService.Update(cur)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Update Category",
		Data:   category,
	}

	helper.WriteEncodeJson(wr, webResponse)
}

func (controller *CategoryControllerImpl) Delete(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	param := params.ByName("id")
	id, err := strconv.Atoi(param)
	helper.PanicHandler(err)

	controller.CategoryService.Delete(id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Delete Category",
	}

	helper.WriteEncodeJson(wr, webResponse)
}
