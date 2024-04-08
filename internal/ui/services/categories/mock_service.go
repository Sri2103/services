package categories_service

import products_templ "github.com/Sri2103/services/internal/ui/views/products"

type mockService struct{}

// GetCategories implements CategoryService.
func (m *mockService) GetCategories() ([]products_templ.ProductCategory, error) {
	return []products_templ.ProductCategory{
		{
			Category: "Fruits",
			Url:      " /products/category/fruits",
			Img:      "https://images.unsplash.com/photo-1579613832125-5d34a13ffe2a?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YXBwbGV8ZW58MHx8MHx8fDA%3D",
		},
		{
			Category: "Food",
			Url:      " /products/category/food",
			Img:      "https://images.unsplash.com/photo-1608039755401-742074f0548d?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8Y2hpY2tlbiUyMHdpbmdzfGVufDB8fDB8fHww",
		},
		{
			Category: "Vegetables",
			Url:      " /products/category/vegetables",
			Img:      "https://plus.unsplash.com/premium_photo-1671013032586-3e9a5c49ce3c?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTN8fG9yYW5nZXxlbnwwfHwwfHx8MA%3D%3D",
		},
		{
			Category: "Drinks",
			Url:      " /products/category/drinks",
			Img:      "https://images.unsplash.com/photo-1608039755401-742074f0548d?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8Y2hpY2tlbiUyMHdpbmdzfGVufDB8fDB8fHww",
		},
		{
			Category: "Dairy",
			Url:      " /products/category/dairy",
			Img:      "https://plus.unsplash.com/premium_photo-1671013032586-3e9a5c49ce3c?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTN8fG9yYW5nZXxlbnwwfHwwfHx8MA%3D%3D",
		},
	}, nil
}

func NewMockService() CategoryService {
	return &mockService{} // Assuming mockService implements CategoryService interface
}
