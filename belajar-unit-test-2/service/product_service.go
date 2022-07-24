package service

import (
	"belajar-unit-test-2/repository"
	"fmt"
	"time"
)

type ProductService interface {
	IsProductReservable(id int) (bool, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

// Constructor
func NewProductService(repository repository.ProductRepository) *productService {
	return &productService{
		productRepository: repository,
	}
}

func (this *productService) IsProductReservable(id int) (bool, error) {
	// Get product information from database
	product, err := this.productRepository.GetProduct(id)
	if err != nil {
		return false, fmt.Errorf("Failed to get product details: %w", err)
	}

	if product == nil {
		return false, fmt.Errorf("Product not found for id %v", id)
	}

	// Only products added more than 1 year ago to the catalog can be reserved
	return product.CreatedAt.Before(time.Now().AddDate(-1, 0, 0)), nil
}
