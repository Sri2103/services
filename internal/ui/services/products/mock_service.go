package product_service

import "github.com/Sri2103/services/internal/ui/views/components"

type mockProductService struct {
}

func NewMockService() ProductService {
	return &mockProductService{}
}

func (m *mockProductService) GetProducts() ([]components.Product, error) {
	pr := []components.Product{
		{ProductId: "Id_1", ProductName: "Apple", ProductPrice: "$5", ProductColor: []string{"Red", "green"}, ProductCategory: "Fruits", ProductImages: []string{"https://images.unsplash.com/photo-1579613832125-5d34a13ffe2a?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YXBwbGV8ZW58MHx8MHx8fDA%3D"}},
		{ProductId: "Id_2", ProductName: "Orange", ProductPrice: "$3", ProductColor: []string{"Orange"}, ProductCategory: "Fruits", ProductImages: []string{"https://plus.unsplash.com/premium_photo-1671013032586-3e9a5c49ce3c?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTN8fG9yYW5nZXxlbnwwfHwwfHx8MA%3D%3D"}},
		{ProductId: "Id_3", ProductName: "Chicken Wings", ProductPrice: "$3", ProductColor: []string{"Violet"}, ProductCategory: "Food", ProductImages: []string{"https://images.unsplash.com/photo-1608039755401-742074f0548d?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8Y2hpY2tlbiUyMHdpbmdzfGVufDB8fDB8fHww"}},
	}
	return pr, nil
}

func (m *mockProductService) AddProduct(product components.Product) error {

	return nil
}

func (m *mockProductService) UpdateProduct(id string, product components.Product) (components.Product, error) {
	var updatedProduct components.Product

	if updatedProduct.ProductName == "" {
		updatedProduct = product
	}
	return updatedProduct, nil
}
