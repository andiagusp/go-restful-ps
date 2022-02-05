package exception

import (
	"golang-restful/helper"
	"golang-restful/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(wr http.ResponseWriter, req *http.Request, err interface{}) {
	if notFoundError(wr, req, err) {
		return
	} else if validationError(wr, req, err) {
		return
	}

	internalServerError(wr, req, err)
}

func notFoundError(wr http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		wr.Header().Set("Content-Type", "application/json")
		wr.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Data Not Found",
			Data:   exception.Error,
		}
		helper.WriteEncodeJson(wr, webResponse)
		return true
	}
	return false
}

func validationError(wr http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		wr.Header().Set("Content-Type", "application/json")
		wr.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Bad Request",
			Data:   exception.Error(),
		}
		helper.WriteEncodeJson(wr, webResponse)
		return true
	}

	return false
}

func internalServerError(wr http.ResponseWriter, req *http.Request, err interface{}) {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal server error",
	}

	helper.WriteEncodeJson(wr, webResponse)
}
