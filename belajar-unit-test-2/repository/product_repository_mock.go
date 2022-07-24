package repository

import (
	"belajar-unit-test-2/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (this *ProductRepositoryMock) GetProduct(id int) (*entity.Product, error) {
	args := this.Mock.Called(id)
	return args.Get(0).(*entity.Product), args.Error(1)
}
