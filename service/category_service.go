package service

import (
	"errors"
	"learn-unit-test/entity"
	"learn-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) FindByID(id string) (*entity.Category, error) {
	category := service.Repository.FindByID(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
