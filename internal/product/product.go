package product

import (
	"context"
	"sle/helpers"
	"time"
)

type Product struct {
	ID          string    `json:"p_id"`
	Name        string    `json:"p_name"`
	Category    string    `json:"p_category"`
	Brand       string    `json:"p_brand"`
	Status      bool      `json:"p_status"`
	Description string    `json:"p_description"`
	Quantity    int       `json:"p_quantity"`
	Image       string    `json:"p_image"`
	Price       int       `json:"p_price"`
	SID         string    `json:"s_id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func NewProduct(name, category, brand, description, image, sid string, quantity, price int) *Product {
	current_time, _ := helpers.GetTime()

	return &Product{
		Name:        name,
		Category:    category,
		Brand:       brand,
		Status:      true,
		Description: description,
		Quantity:    quantity,
		Image:       image,
		Price:       price,
		SID:         sid,
		Created_at:  current_time,
		Updated_at:  current_time,
	}
}

type ProductReq struct {
	Name        string `json:"p_name"`
	Category    string `json:"p_category"`
	Brand       string `json:"p_brand"`
	Description string `json:"p_description"`
	Quantity    int    `json:"p_quantity"`
	Image       string `json:"p_image"`
	Price       int    `json:"p_price"`
	SID         string `json:"s_id"`
}

type ProductRes struct {
	ID           string    `json:"p_id"`
	Name         string    `json:"p_name"`
	Category     string    `json:"p_category"`
	Brand        string    `json:"p_brand"`
	Status       bool      `json:"p_status"`
	Description  string    `json:"p_description"`
	Quantity     int       `json:"p_quantity"`
	Image        string    `json:"p_image"`
	Price        int       `json:"p_price"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	SID          string    `json:"s_id"`
	S_First_Name string    `json:"s_first_name"`
	S_Last_Name  string    `json:"s_last_name"`
	S_Phone      string    `json:"s_phone"`
}

type Repository interface {
	CreateProduct(ctx context.Context, product *Product, seller_id string) (*Product, error)
	GetProductByID(ctx context.Context, pid string) (*ProductRes, error)
	GetAllProducts(ctx context.Context, page int, recordPerPage int) (*[]ProductRes, error)
	GetAllProductsBySellerAndCategory(ctx context.Context, s_id string, category string, page int, recordPerPage int) (*[]ProductRes, error)
	DeleteProduct(ctx context.Context, pid string) error
	UpdateProductDetails(ctx context.Context, pid string, req ProductReq) error
	UpdateStatus(ctx context.Context, pid string) error
	SearchProduct(ctx context.Context, str string, page int, recordPerPage int) (*[]ProductRes, error)
}

type Service interface {
	CreateProduct(ctx context.Context, product *ProductReq, seller_id string) (*Product, error)
	GetProductByID(ctx context.Context, pid string) (*ProductRes, error)
	GetAllProducts(ctx context.Context, page int, recordPerPage int) (*[]ProductRes, error)
	GetAllProductsBySellerAndCategory(ctx context.Context, s_id string, category string, page int, recordPerPage int) (*[]ProductRes, error)
	DeleteProduct(ctx context.Context, pid string) error
	UpdateProductDetails(ctx context.Context, pid string, req ProductReq) error
	UpdateStatus(ctx context.Context, pid string) error
	SearchProduct(ctx context.Context, str string, page int, recordPerPage int) (*[]ProductRes, error)
}
