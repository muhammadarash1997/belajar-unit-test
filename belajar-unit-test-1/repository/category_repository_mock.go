package repository

import (
	"belajar-unit-test-1/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (this *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := this.Mock.Called(id)
	if arguments.Get(0) == nil {	// 0 di sini maksudnya mengambil data pertama
		return nil
	}
	category := arguments.Get(0).(entity.Category)
	return &category
}
