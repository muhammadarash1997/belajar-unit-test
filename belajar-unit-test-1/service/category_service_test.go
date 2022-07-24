package service

import (
	"belajar-unit-test-1/entity"
	"belajar-unit-test-1/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// Program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}

	// Program mock
	categoryRepository.Mock.On("FindById", "2").Return(category)	// 2 adalah argument untuk method FindById

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}

// Dalam test ini kita ingin mengetest fungsi Get pada pada layer Service yang mana fungsi Get tsb
// harus memanggil fungsi FindById pada layer Repository dan yang jadi masalah adalah fungsi FindById
// tsb adalah fungsi dari sebuah objek yang memiliki objek third-party maka untuk menyelesaikan
// masalah tsb kita perlu menggunakan Mock
