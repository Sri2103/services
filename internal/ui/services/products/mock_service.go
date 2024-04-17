package product_service

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/Sri2103/services/internal/ui/views/components"
)

type mockProductService struct {
	data []components.Product
}

// AddCategory implements ProductService.
func (m *mockProductService) AddCategory(category string) error {
	panic("unimplemented")
}

func NewMockService() ProductService {
	return &mockProductService{
		data: GenerateDummyData(),
	}
}

func GenerateDummyData() []components.Product {
	var products []components.Product

	// Fruits
	for i := 1; i <= 20; i++ {
		product := components.Product{
			ProductId:          fmt.Sprintf("FR-%d", i),
			ProductName:        fmt.Sprintf("Fruit%d", i),
			ProductColor:       []string{"Red", "Green", "Yellow"},
			ProductCategory:    "fruits",
			ProductPrice:       fmt.Sprintf("%.2f", rand.Float64()*5.0),
			ProductDescription: fmt.Sprintf("Description for Fruit%d", i),
			ProductImages:      []string{"https://images.unsplash.com/photo-1579613832125-5d34a13ffe2a?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YXBwbGV8ZW58MHx8MHx8fDA%3D"},
		}
		products = append(products, product)
	}

	// Food
	for i := 1; i <= 20; i++ {
		product := components.Product{
			ProductId:          fmt.Sprintf("FD-%d", i),
			ProductName:        fmt.Sprintf("Food%d", i),
			ProductColor:       []string{"Brown", "White", "Yellow"},
			ProductCategory:    "food",
			ProductPrice:       fmt.Sprintf("%.2f", rand.Float64()*5.0),
			ProductDescription: fmt.Sprintf("Description for Food%d", i),
			ProductImages:      []string{"https://images.unsplash.com/photo-1579613832125-5d34a13ffe2a?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YXBwbGV8ZW58MHx8MHx8fDA%3D"},
		}
		products = append(products, product)
	}

	// Vegetables
	for i := 1; i <= 20; i++ {
		product := components.Product{
			ProductId:          fmt.Sprintf("VG-%d", i),
			ProductName:        fmt.Sprintf("Vegetable%d", i),
			ProductColor:       []string{"Green", "Yellow", "Purple"},
			ProductCategory:    "vegetables",
			ProductPrice:       fmt.Sprintf("%.2f", rand.Float64()*5.0),
			ProductDescription: fmt.Sprintf("Description for Vegetable%d", i),
			ProductImages:      []string{"https://images.unsplash.com/photo-1579613832125-5d34a13ffe2a?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YXBwbGV8ZW58MHx8MHx8fDA%3D"},
		}
		products = append(products, product)
	}

	// Drinks
	for i := 1; i <= 20; i++ {
		product := components.Product{
			ProductId:          fmt.Sprintf("DR-%d", i),
			ProductName:        fmt.Sprintf("Drink%d", i),
			ProductColor:       []string{"Clear", "Brown", "Red"},
			ProductCategory:    "drinks",
			ProductPrice:       fmt.Sprintf("%.2f", rand.Float64()*5.0),
			ProductDescription: fmt.Sprintf("Description for Drink%d", i),
			ProductImages:      []string{"https://images.unsplash.com/photo-1579613832125-5d34a13ffe2a?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YXBwbGV8ZW58MHx8MHx8fDA%3D"},
		}
		products = append(products, product)
	}

	// Dairy
	for i := 1; i <= 20; i++ {
		product := components.Product{
			ProductId:          fmt.Sprintf("DY-%d", i),
			ProductName:        fmt.Sprintf("Dairy%d", i),
			ProductColor:       []string{"White", "Yellow"},
			ProductCategory:    "dairy",
			ProductPrice:       fmt.Sprintf("%.2f", rand.Float64()*5.0),
			ProductDescription: fmt.Sprintf("Description for Dairy%d", i),
			ProductImages:      []string{"https://images.unsplash.com/photo-1579613832125-5d34a13ffe2a?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YXBwbGV8ZW58MHx8MHx8fDA%3D"},
		}
		products = append(products, product)
	}

	return products
}

func (m *mockProductService) GetProducts() ([]components.Product, error) {
	pr := []components.Product{
		{ProductId: "Id_1", ProductName: "Apple", ProductPrice: "$5", ProductColor: []string{"Red", "green"}, ProductCategory: "Fruits", ProductImages: []string{"https://images.unsplash.com/photo-1579613832125-5d34a13ffe2a?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8YXBwbGV8ZW58MHx8MHx8fDA%3D"}},
		{ProductId: "Id_2", ProductName: "Orange", ProductPrice: "$3", ProductColor: []string{"Orange"}, ProductCategory: "Fruits", ProductImages: []string{"https://plus.unsplash.com/premium_photo-1671013032586-3e9a5c49ce3c?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTN8fG9yYW5nZXxlbnwwfHwwfHx8MA%3D%3D"}},
		{ProductId: "Id_3", ProductName: "Chicken Wings", ProductPrice: "$3", ProductColor: []string{"Violet"}, ProductCategory: "Food", ProductImages: []string{"https://images.unsplash.com/photo-1608039755401-742074f0548d?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8Y2hpY2tlbiUyMHdpbmdzfGVufDB8fDB8fHww"}},
	}
	return pr, nil
}

// AddProduct
func (m *mockProductService) AddProduct(product components.Product) error {
	fmt.Println(product, "Add Product from the Add Product API called")

	return nil
}

func (m *mockProductService) UpdateProduct(id string, product components.Product) (components.Product, error) {
	var updatedProduct components.Product

	if updatedProduct.ProductName == "" {
		updatedProduct = product
	}
	return updatedProduct, nil
}

func (s *mockProductService) GetProductsByCategory(category string, pageNumber int, pageSize int, sort string) ([]components.Product, int, error) {
	// get category from the data with category, pagination, pagesize
	var products []components.Product
	startIndex := (pageNumber-1)*pageSize + 1
	endIndex := startIndex + pageSize
	categoryProducts := s.getProductsByCategory(category, sort)
	if endIndex > len(categoryProducts) {
		endIndex = len(categoryProducts)
	}
	for i := startIndex; i < endIndex; i++ {
		products = append(products, categoryProducts[i-1])
	}

	numberOfResultPages := int(math.Ceil(float64(len(categoryProducts)) / float64(pageSize)))
	return products, numberOfResultPages, nil
}

func (s *mockProductService) getProductsByCategory(category string, sort string) []components.Product {
	var categoryProducts []components.Product
	for _, product := range s.data {
		if product.ProductCategory == category {
			categoryProducts = append(categoryProducts, product)
		}
	}
	if sort != "" {
		categoryProducts = s.SortProducts(categoryProducts, sort)
	}
	return categoryProducts
}

// sort the products based on price asc or desc
func (s *mockProductService) SortProducts(products []components.Product, sortString string) []components.Product {

	if sortString == "desc" {
		sort.Slice(products, func(i, j int) bool {
			return products[i].ProductPrice > products[j].ProductPrice
		})
	} else if sortString == "asc" {
		sort.Slice(products, func(i, j int) bool {
			return products[i].ProductPrice < products[j].ProductPrice
		})
	}
	return products
}

// get ProductById
func (s *mockProductService) GetProductById(id string) (components.Product, error) {

	for _, product := range s.data {
		if product.ProductId == id {
			return product, nil
		}
	}
	return components.Product{}, errors.New("product not found")
}
