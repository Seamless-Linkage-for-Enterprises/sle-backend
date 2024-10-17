package product

import (
	"context"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewProductService(r Repository) Service {
	return &service{Repository: r, timeout: time.Duration(100) * time.Second}
}

func (s *service) CreateProduct(ctx context.Context, reqproduct *ProductReq, seller_id string) (*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	product := NewProduct(reqproduct.Name, reqproduct.Category, reqproduct.Brand, reqproduct.Description, reqproduct.Image, seller_id, reqproduct.Quantity, reqproduct.Price)

	return s.Repository.CreateProduct(ctx, product, seller_id)
}
func (s *service) GetProductByID(ctx context.Context, pid string) (*ProductRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.GetProductByID(ctx, pid)
}

func (s *service) GetAllProducts(ctx context.Context, page int, recordPerPage int) (*[]ProductRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.GetAllProducts(ctx, page, recordPerPage)
}

func (s *service) GetAllProductsBySellerAndCategory(ctx context.Context, s_id, category string, page int, recordPerPage int) (*[]ProductRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.GetAllProductsBySellerAndCategory(ctx, s_id, category, page, recordPerPage)
}

func (s *service) DeleteProduct(ctx context.Context, pid string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.DeleteProduct(ctx, pid)
}

func (s *service) UpdateProductDetails(ctx context.Context, pid string, req ProductReq) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.UpdateProductDetails(ctx, pid, req)
}

func (s *service) UpdateStatus(ctx context.Context, pid string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.UpdateStatus(ctx, pid)
}

func (s *service) SearchProduct(ctx context.Context, str string, page int, recordPerPage int) (*[]ProductRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.SearchProduct(ctx, str, page, recordPerPage)
}
