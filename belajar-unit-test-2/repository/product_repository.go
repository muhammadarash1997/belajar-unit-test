package repository

import "belajar-unit-test-2/entity"

type ProductRepository interface {
	GetProduct(id int) (*entity.Product, error)
}
