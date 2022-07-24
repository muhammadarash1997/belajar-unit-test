package service

import (
	"belajar-unit-test-2/entity"
	"belajar-unit-test-2/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewProductService(t *testing.T) {
	assertions := require.New(t)

	// Register test mock and create service with passing mock
	productRepositoryMock := repository.ProductRepositoryMock{} // Ignore the mock for now

	// Test
	productServiceObj := NewProductService(&productRepositoryMock)

	// Check
	assertions.Implements((*ProductService)(nil), new(productService), "Product Service Implementation does not honor service definition")  // <-- Check bahwa struct productService harus mengimplement interface ProductService, jika tidak maka testing akan Fail dan menampilkan pesan "Product Service Implementation does not honor service definition"
	assertions.NotNil(t, productServiceObj, "Product Service not initialized") // <-- Check bahwa productServiceObj terbuat dan tidak nil, jika tidak maka testing akan Fail dan menampilkan pesan "Product Service not initialized"
	assertions.NotNil(t, productServiceObj.productRepository, "Product Service dependency not initialized") // <-- Check bahwa productRepository terbuat dan tidak nil, jika tidak maka testing akan Fail dan menampilkan pesan "Product Service dependency not initialized"
}

func TestProductService_IsProductReservable(t *testing.T) {
	assertions := require.New(t)

	// Register test mock and create service with passing mock
	productRepositoryMock := repository.ProductRepositoryMock{}
	productService := NewProductService(&productRepositoryMock)

	// Program mock
	productRepositoryMock.Mock.On("GetProduct", 1).Return(&entity.Product{
		Id:          1,
		Description: "Product created 2 years ago",
		CreatedAt:   time.Now().AddDate(-2, 0, 0),
	}, nil)

	// Program mock
	productRepositoryMock.Mock.On("GetProduct", 2).Return(&entity.Product{
		Id:          2,
		Description: "Product recently created",
		CreatedAt:   time.Now(),
	}, nil)

	testDataSet := map[int]bool{
		1: true,
		2: false,
	}

	for productId, expectedResult := range testDataSet {
		// Test
		reservable, err := productService.IsProductReservable(productId)

		// Check
		assertions.NoErrorf(err, "Failed to check if product %v is reservable: %s", productId, err)
		assertions.Equalf(expectedResult, reservable, "Got wrong reservable info for product id %v", productId)
	}
}

func TestProductService_IsProductReservable_NotFound(t *testing.T) {
	assertions := require.New(t)

	// Register test mock and create service with passing mock
	productRepositoryMock := repository.ProductRepositoryMock{}
	productService := NewProductService(&productRepositoryMock)

	// Program mock
	productRepositoryMock.Mock.On("GetProduct", 1).Return((*entity.Product)(nil), nil)

	// Test
	_, err := productService.IsProductReservable(1)

	// Check
	assertions.NotNil(t, err)
}
