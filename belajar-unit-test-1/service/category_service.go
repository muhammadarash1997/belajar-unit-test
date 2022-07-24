package service

import (
	"belajar-unit-test-1/entity"
	"belajar-unit-test-1/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (this *CategoryService) Get(id string) (*entity.Category, error) {
	category := this.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found")
	}
	return category, nil
}
