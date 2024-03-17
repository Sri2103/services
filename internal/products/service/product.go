package productImplementation

import (
	"context"

	repo "github.com/Sri2103/services/internal/products/repository"
	"github.com/Sri2103/services/pkg/ent"
	product_pb "github.com/Sri2103/services/pkg/rpc/product"
	"github.com/google/uuid"
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
	productID, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, err
	}
	product, err := p.repo.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}
	gr.Product = &product_pb.Product{
		Id:          product.ID.String(),
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
			Id:          v.ID.String(),
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
		})

	}
	return &lr, nil
}

func (p *productImpl) CreateProduct(ctx context.Context, r *product_pb.CreateProductRequest) (*product_pb.CreateProductResponse, error) {
	var res product_pb.CreateProductResponse
	if r.Product.Id == "" {
		r.Product.Id = uuid.NewString()
	}
	id, err := uuid.Parse(r.Product.Id)
	if err != nil {
		return nil, err
	}

	pr := ent.Product{
		ID:          id,
		Name:        r.Product.Name,
		Description: r.Product.Description,
		Price:       r.Product.Price,
	}
	pC, err := p.repo.CreateProduct(ctx, &pr)
	if err != nil {
		return nil, err
	}

	res.Product = &product_pb.Product{
		Id:          pC.ID.String(),
		Name:        pC.Name,
		Description: pC.Description,
		Price:       pC.Price,
	}

	return &res, nil
}
