package productImplementation

import (
	"context"
	"fmt"
	"math"

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
		Price:       float32(product.Price),
		Images:      product.Images,
		Colors:      product.Color,
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
			Price:       float32(v.Price),
			Images:      v.Images,
			Colors:      v.Color,
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
		Price:       float64(r.Product.Price),
		Images:      r.Product.Images,
		Color:       r.Product.Colors,
	}
	category := ent.Category{
		ID: uuid.MustParse(r.Category.Id),
	}
	pC, err := p.repo.CreateProduct(ctx, &pr, &category) // Added nil as the third argument
	if err != nil {
		return nil, err
	}

	res.Product = &product_pb.Product{
		Id:          pC.ID.String(),
		Name:        pC.Name,
		Description: pC.Description,
		Price:       float32(pC.Price),
		Images:      pC.Images,
		Colors:      pC.Color,
	}

	return &res, nil
}

// update products
func (p *productImpl) UpdateProduct(ctx context.Context, r *product_pb.UpdateProductRequest) (*product_pb.UpdateProductResponse, error) {
	var res product_pb.UpdateProductResponse
	id, err := uuid.Parse(r.Product.Id)
	if err != nil {
		return nil, err
	}
	pr := ent.Product{
		ID:          id,
		Name:        r.Product.Name,
		Description: r.Product.Description,
		Price:       float64(r.Product.Price),
		Images:      r.Product.Images,
		Color:       r.Product.Colors,
	}
	pC, err := p.repo.UpdateProduct(ctx, &pr)
	if err != nil {
		return nil, err
	}
	res.Product = &product_pb.Product{
		Id:          pC.ID.String(),
		Name:        pC.Name,
		Description: pC.Description,
		Price:       float32(pC.Price),
		Images:      pC.Images,
		Colors:      pC.Color,
	}
	return &res, nil
}

// create category
func (p *productImpl) CreateCategory(ctx context.Context, r *product_pb.CreateCategoryRequest) (*product_pb.CreateCategoryResponse, error) {
	var res product_pb.CreateCategoryResponse
	cat := ent.Category{
		Name: r.Name,
	}
	pC, err := p.repo.CreateCategory(ctx, &cat)
	if err != nil {
		return nil, err
	}
	res.Id = pC.ID.String()
	res.Name = pC.Name
	return &res, nil
}

// Get All Categories
func (p *productImpl) GetAllCategories(ctx context.Context, r *product_pb.GetAllCategoriesRequest) (*product_pb.GetAllCategoriesResponse, error) {
	var ctg []*product_pb.Category
	categories, err := p.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range categories {
		ctg = append(ctg, &product_pb.Category{
			Id:   v.ID.String(),
			Name: v.Name,
		})
	}
	return &product_pb.GetAllCategoriesResponse{
		Categories: ctg,
	}, nil
}

// Get productsByCategories
func (p *productImpl) GetProductsByCategory(ctx context.Context, r *product_pb.GetProductsByCategoryRequest) (*product_pb.GetProductsByCategoryResponse, error) {
	fmt.Println(r.CategoryId, r.PageNumber, r.ResultsPerPage, r.Sort, "ciD", "pn", "rp", "s")
	var pr []*product_pb.Product
	products, totalProducts, err := p.repo.GetCategorizedProducts(ctx, uuid.MustParse(r.CategoryId), int(r.PageNumber), int(r.ResultsPerPage), r.Sort)
	if err != nil {
		fmt.Println("err from the getProductsByCategory", err)
		return nil, err
	}

	for _, v := range products {
		pr = append(pr, &product_pb.Product{
			Id:          v.ID.String(),
			Name:        v.Name,
			Description: v.Description,
			Price:       float32(v.Price),
			Images:      v.Images,
			Colors:      v.Color,
		})
	}
	return &product_pb.GetProductsByCategoryResponse{
		Products:     pr,
		TotalResults: int32(totalProducts),
		TotalPages:   int32(math.Ceil(float64(totalProducts) / float64(r.ResultsPerPage))),
	}, nil
}

// Get Category By ID
func (p *productImpl) GetCategory(ctx context.Context, r *product_pb.GetCategoryRequest) (*product_pb.Category, error) {
	var ctg product_pb.Category
	category, err := p.repo.GetCategory(ctx, uuid.MustParse(r.Id))
	if err != nil {
		return nil, err
	}
	ctg.Id = category.ID.String()
	ctg.Name = category.Name
	return &ctg, nil
}
