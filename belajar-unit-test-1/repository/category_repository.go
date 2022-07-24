package repository

import (
	"belajar-unit-test-1/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (this *categoryRepository) FindById(id string) *entity.Category {
	category := &entity.Category{}
	err := this.db.First(&category, "id = ?", id).Error
	if err != nil {
		return nil
	}

	return category
}
