package repository

import (
	"database/sql"
	"golang-restful/model/domain"
)

type CategoryRepository interface {
	Save(tx *sql.Tx, category domain.Category) domain.Category
	FindAll(tx *sql.Tx) []domain.Category
	Update(tx *sql.Tx, category domain.Category) domain.Category
	Delete(tx *sql.Tx, id int)
	FindById(tx *sql.Tx, id int) (domain.Category, error)
}
