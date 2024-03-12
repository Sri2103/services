package productImplementation

import (
	"context"

	repo "github.com/Sri2103/services/internal/products/repository"
	product_pb "github.com/Sri2103/services/pkg/rpc/product"
)

type productImpl struct {
	repo repo.Repo
	product_pb.UnimplementedProductServiceServer
}

func New(repo repo.Repo) product_pb.ProductServiceServer {
	return &productImpl{
		repo: repo,
	}
}

func (p *productImpl) GetProduct(ctx context.Context, r *product_pb.GetProductRequest) (*product_pb.GetProductResponse, error) {
	var gr product_pb.GetProductResponse
	product, err := p.repo.GetProduct(ctx, int(r.Id))
	if err != nil {
		return nil, err
	}
	gr.Product = &product_pb.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
	return &gr, nil
}

func (p *productImpl) ListProducts(ctx context.Context, _ *product_pb.ListProductsRequest) (*product_pb.ListProductsResponse, error) {
	var lr product_pb.ListProductsResponse
	products, err := p.repo.GetProductList(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range products {
		lr.Products = append(lr.Products, &product_pb.Product{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
		})

	}
	return &lr, nil
}
