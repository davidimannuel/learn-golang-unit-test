package service

import (
	"learn-unit-test/entity"
	"learn-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_FindByID_NotFound(t *testing.T) {
	categoryRepository.Mock.On("FindByIDc", "1").Return(nil)
	category, err := categoryService.FindByID("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}
func TestCategoryService_FindByID_Found(t *testing.T) {
	category := entity.Category{
		ID:   "002",
		Name: "Elektronic",
	}
	categoryRepository.Mock.On("FindByID", "002").Return(category)
	result, err := categoryService.FindByID("002")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}
