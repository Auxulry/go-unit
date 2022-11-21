package repository

import "go-unit/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
