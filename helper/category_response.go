package helper

import (
	"golang-restful/model/domain"
	"golang-restful/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryListResponse(category []domain.Category) []web.CategoryResponse {
	var categories []web.CategoryResponse

	for _, data := range category {
		categories = append(categories, ToCategoryResponse(data))
	}

	return categories
}
