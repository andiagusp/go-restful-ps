package repository

import (
	"database/sql"
	"golang-restful/helper"
	"golang-restful/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) VALUES($1) RETURNING id"

	row := tx.QueryRow(SQL, category.Name)
	err := row.Scan(&category.Id)
	helper.PanicHandler(err)

	return category
}

func (repository *CategoryRepositoryImpl) FindAll(tx *sql.Tx) []domain.Category {
	var categories []domain.Category
	SQL := "SELECT id, name FROM category ORDER BY id DESC"

	rows, err := tx.Query(SQL)
	helper.PanicHandler(err)

	for rows.Next() {
		var category domain.Category
		err2 := rows.Scan(&category.Id, &category.Name)
		helper.PanicHandler(err2)
		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) Update(tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = $1 WHERE id = $2"
	_, err := tx.Exec(SQL, category.Name, category.Id)
	helper.PanicHandler(err)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(tx *sql.Tx, id int) {
	SQL := "DELETE FROM category WHERE id = $1"
	_, err := tx.Exec(SQL, id)

	helper.PanicHandler(err)
}

func (repository *CategoryRepositoryImpl) FindById(tx *sql.Tx, id int) (domain.Category, error) {
	var category domain.Category
	SQL := "SELECT id, name FROM category WHERE id = $1"

	row := tx.QueryRow(SQL, id)
	err := row.Scan(&category.Id, &category.Name)

	return category, err
}
