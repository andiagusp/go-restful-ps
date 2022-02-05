package main

import (
	"golang-restful/app"
	"golang-restful/controller"
	"golang-restful/database"
	"golang-restful/helper"
	"golang-restful/middleware"
	"golang-restful/repository"
	"golang-restful/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {
	db := database.NewDB()
	validator := validator.New()
	cr := repository.NewCategoryRepository()
	cs := service.NewCategoryService(db, validator, cr)
	categoryController := controller.NewCategoryController(cs)

	router := app.NewRouter(categoryController)
	middleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: middleware,
	}

	err := server.ListenAndServe()
	helper.PanicHandler(err)

	log.Println("Server Running On Port 5000")
}
