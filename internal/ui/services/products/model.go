package product_service

import (
	"strconv"

	"github.com/Sri2103/services/internal/ui/views/components"
)

type Product struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	Colors      []string `json:"color"`
	Images      []string `json:"images"`
}



func convertToComponentModal(product Product) components.Product {
	return components.Product{
		ProductId:          product.Id,
		ProductName:        product.Name,
		ProductDescription: product.Description,
		ProductPrice:       strconv.FormatFloat(float64(product.Price), 'f', 2, 32),
		ProductColor:       product.Colors,
		ProductImages:      product.Images,
	}
}

func convertToProduct(product components.Product) Product {
	price, _ := strconv.ParseFloat(product.ProductPrice, 32)
	return Product{
		Id:          product.ProductId,
		Name:        product.ProductName,
		Description: product.ProductDescription,
		Price:       float32(price),
		Colors:      product.ProductColor,
		Images:      product.ProductImages,
	}
}

func convertToProducts(products []Product) []components.Product {
	var componentsProducts []components.Product
	for _, product := range products {
		componentsProducts = append(componentsProducts, convertToComponentModal(product))
	}
	return componentsProducts
}
