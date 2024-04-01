package product_service

import "github.com/Sri2103/services/internal/ui/views/components"

type mockProductService struct {
}

func NewMockService() ProductService {
	return &mockProductService{}
}

func (m *mockProductService) GetProducts() ([]components.Product, error) {
	pr := []components.Product{
		{ProductId: "Id_1", ProductName: "Apple", ProductPrice: "$5", ProductColor: []string{"Red", "green"}, ProductCategory: "Fruits"},
		{ProductId: "Id_2", ProductName: "Orange", ProductPrice: "$3", ProductColor: []string{"Orange"}, ProductCategory: "Fruits"},
		{ProductId: "Id_3", ProductName: "Chicken Wings", ProductPrice: "$3", ProductColor: []string{"Violet"}, ProductCategory: "Food"},
	}
	return pr, nil
}

func (m *mockProductService) AddProduct(product components.Product) error {

	return nil
}
