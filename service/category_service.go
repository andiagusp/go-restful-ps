package service

import "golang-restful/model/web"

type CategoryService interface {
	Delete(id int)
	FindAll() []web.CategoryResponse
	FindById(id int) web.CategoryResponse
	Update(category web.CategoryUpdateRequest) web.CategoryResponse
	Create(category web.CategoryCreateRequest) web.CategoryResponse
}
