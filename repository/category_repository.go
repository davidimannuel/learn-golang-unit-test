package repository

import "learn-unit-test/entity"

type CategoryRepository interface {
	FindByID(id string) *entity.Category
}
