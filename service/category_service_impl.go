package service

import (
	"database/sql"
	"golang-restful/exception"
	"golang-restful/helper"
	"golang-restful/model/domain"
	"golang-restful/model/web"
	"golang-restful/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	DB                 *sql.DB
	Validate           *validator.Validate
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(db *sql.DB, validate *validator.Validate, cr repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		DB:                 db,
		Validate:           validate,
		CategoryRepository: cr,
	}
}

func (service *CategoryServiceImpl) Delete(id int) {
	tx, err := service.DB.Begin()
	helper.PanicHandler(err)
	defer helper.CommitOrRollback(tx)

	_, err2 := service.CategoryRepository.FindById(tx, id)
	if err2 != nil {
		exception.NewNotFoundError(err2.Error())
	}
	service.CategoryRepository.Delete(tx, id)
}

func (service *CategoryServiceImpl) FindAll() []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicHandler(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(tx)

	return helper.ToCategoryListResponse(categories)
}

func (service *CategoryServiceImpl) FindById(id int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicHandler(err)
	defer helper.CommitOrRollback(tx)

	category, err2 := service.CategoryRepository.FindById(tx, id)
	if err2 != nil {
		exception.NewNotFoundError(err2.Error())
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicHandler(err)
	tx, err := service.DB.Begin()
	helper.PanicHandler(err)
	defer helper.CommitOrRollback(tx)

	category, err2 := service.CategoryRepository.FindById(tx, request.Id)
	if err2 != nil {
		exception.NewNotFoundError(err2.Error())
	}
	category.Name = request.Name
	category = service.CategoryRepository.Update(tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Create(request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicHandler(err)
	tx, err := service.DB.Begin()
	helper.PanicHandler(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(tx, category)

	return helper.ToCategoryResponse(category)
}
