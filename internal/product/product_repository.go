package product

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

// * create product
func (r *repository) CreateProduct(ctx context.Context, product *Product, seller_id string) (*Product, error) {

	query := `INSERT INTO products (p_name,p_category,p_brand,p_status,p_description,p_quantity,p_image,p_price,s_id)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING p_id,created_at,updated_at`

	err := r.db.QueryRow(ctx, query, product.Name, product.Category, product.Brand, product.Status, product.Description, product.Quantity, product.Image, product.Price, seller_id).Scan(&product.ID, &product.Created_at, &product.Updated_at)

	if err != nil {
		return nil, err
	}

	if product.ID == "" {
		return nil, errors.New("failed to insert record")
	}

	product.Status = true

	return product, nil
}

// * get all products
func (r *repository) GetAllProducts(ctx context.Context, page int, recordPerPage int) (*[]ProductRes, error) {
	offset := (page - 1) * recordPerPage
	log.Println(offset, page, recordPerPage)
	query := `
SELECT 
    p.p_id,
    p.p_name,
    p.p_category,
    p.p_brand,
    p.p_status,
    p.p_description,
    p.p_quantity,
    p.p_image,
    p.p_price,
    p.created_at,
    p.updated_at,
	s.s_id,
    s.s_phone,
	s.s_first_name,
	s.s_last_name
FROM 
    products p
JOIN 
    sellers s ON p.s_id = s.s_id
ORDER BY p.created_at DESC
OFFSET $1 
LIMIT $2
	`
	var products []ProductRes
	row, err := r.db.Query(ctx, query, offset, recordPerPage)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var product ProductRes
		if err := row.Scan(
			&product.ID,
			&product.Name,
			&product.Category,
			&product.Brand,
			&product.Status,
			&product.Description,
			&product.Quantity,
			&product.Image,
			&product.Price,
			&product.Created_at,
			&product.Updated_at,
			&product.SID,
			&product.S_Phone,
			&product.S_First_Name,
			&product.S_Last_Name,
		); err != nil {
			log.Println(err.Error())
		}
		products = append(products, product)
	}

	return &products, nil
}

// * get product by id
func (r *repository) GetProductByID(ctx context.Context, pid string) (*ProductRes, error) {
	query := `
SELECT 
    p.p_id,
    p.p_name,
    p.p_category,
    p.p_brand,
    p.p_status,
    p.p_description,
    p.p_quantity,
    p.p_image,
    p.p_price,
    p.created_at,
    p.updated_at,
	s.s_id,
    s.s_phone,
	s.s_first_name,
	s.s_last_name
FROM 
    products p
JOIN 
    sellers s ON p.s_id = s.s_id
WHERE 
    p.p_id = $1;

	`
	var product ProductRes

	if err := r.db.QueryRow(ctx, query, pid).Scan(
		&product.ID,
		&product.Name,
		&product.Category,
		&product.Brand,
		&product.Status,
		&product.Description,
		&product.Quantity,
		&product.Image,
		&product.Price,
		&product.Created_at,
		&product.Updated_at,
		&product.SID,
		&product.S_Phone,
		&product.S_First_Name,
		&product.S_Last_Name,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("product doesn't exist with the provided p_id")
		}
		return nil, err
	}

	return &product, nil
}

// * update product
func (r *repository) UpdateProductDetails(ctx context.Context, pid string, req ProductReq) error {

	// check seller exist with p_id
	_, err := r.GetProductByID(ctx, pid)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("product doesn't exist with the provided p_id")
		}
		return err
	}

	query := `
	UPDATE products 
	SET p_name=$1,p_category=$2,p_brand=$3,p_description=$4,p_quantity=$5,p_image=$6,p_price=$7
	WHERE p_id=$8
	`

	// update product password
	if _, err := r.db.Exec(ctx, query, req.Name, req.Category, req.Brand, req.Description, req.Quantity, req.Image, req.Price, pid); err != nil {
		return err
	}

	return nil
}

// * delete product
func (r *repository) DeleteProduct(ctx context.Context, pid string) error {
	query := `
	DELETE FROM products WHERE p_id=$1
	`

	_, err := r.db.Exec(ctx, query, pid)
	if err != nil {
		return err
	}
	return nil
}

// * change status
func (r *repository) UpdateStatus(ctx context.Context, pid string) error {
	status := false
	// check seller exist with p_id
	res, err := r.GetProductByID(ctx, pid)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("product doesn't exist with the provided p_id")
		}
		return err
	}

	if res.Status {
		status = false
	} else {
		status = true
	}

	query := `
	UPDATE products 
	SET p_status=$1
	WHERE p_id=$2
	`

	// update product password
	if _, err := r.db.Exec(ctx, query, status, pid); err != nil {
		return err
	}

	return nil
}

// * search product by product name
func (r *repository) SearchProduct(ctx context.Context, str string, page int, recordPerPage int) (*[]ProductRes, error) {
	offset := (page - 1) * recordPerPage
	query := `
SELECT 
    p.p_id,
    p.p_name,
    p.p_category,
    p.p_brand,
    p.p_status,
    p.p_description,
    p.p_quantity,
    p.p_image,
    p.p_price,
    p.created_at,
    p.updated_at,
    s.s_id,
    s.s_phone,
    s.s_first_name,
    s.s_last_name
FROM 
    products p
JOIN 
    sellers s ON p.s_id = s.s_id
WHERE 
    p.p_name ILIKE '%' || $1 || '%' -- ILIKE for case-insensitive search
ORDER BY p.created_at DESC
OFFSET $2
LIMIT $3;
	`

	var products []ProductRes
	row, err := r.db.Query(ctx, query, str, offset, recordPerPage)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var product ProductRes
		if err := row.Scan(
			&product.ID,
			&product.Name,
			&product.Category,
			&product.Brand,
			&product.Status,
			&product.Description,
			&product.Quantity,
			&product.Image,
			&product.Price,
			&product.Created_at,
			&product.Updated_at,
			&product.SID,
			&product.S_Phone,
			&product.S_First_Name,
			&product.S_Last_Name,
		); err != nil {
			log.Println(err.Error())
		}
		products = append(products, product)
	}

	return &products, nil
}

// * products by category & s_id
func (r *repository) GetAllProductsBySellerAndCategory(ctx context.Context, s_id string, category string, page int, recordPerPage int) (*[]ProductRes, error) {
	offset := (page - 1) * recordPerPage
	log.Println(s_id, category, ">>", offset, page, recordPerPage)

	// Base query
	query := `
SELECT 
    p.p_id,
    p.p_name,
    p.p_category,
    p.p_brand,
    p.p_status,
    p.p_description,
    p.p_quantity,
    p.p_image,
    p.p_price,
    p.created_at,
    p.updated_at,
    s.s_id,
    s.s_phone,
    s.s_first_name,
    s.s_last_name
FROM 
    products p
JOIN 
    sellers s ON p.s_id = s.s_id
WHERE 1=1
`

	// Slice to hold query parameters
	var args []interface{}
	var argIndex int

	// Add s_id condition if provided
	if s_id != "" {
		argIndex++
		query += fmt.Sprintf(" AND p.s_id = $%d", argIndex)
		args = append(args, s_id)
	}

	// Add category condition if provided
	if category != "" {
		argIndex++
		query += fmt.Sprintf(" AND p.p_category = $%d", argIndex)
		args = append(args, category)
	}

	// Add pagination
	query += fmt.Sprintf(" ORDER BY p.created_at DESC OFFSET $%d LIMIT $%d", argIndex+1, argIndex+2)
	args = append(args, offset, recordPerPage)

	// Execute the query
	var products []ProductRes
	row, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var product ProductRes
		if err := row.Scan(
			&product.ID,
			&product.Name,
			&product.Category,
			&product.Brand,
			&product.Status,
			&product.Description,
			&product.Quantity,
			&product.Image,
			&product.Price,
			&product.Created_at,
			&product.Updated_at,
			&product.SID,
			&product.S_Phone,
			&product.S_First_Name,
			&product.S_Last_Name,
		); err != nil {
			log.Println(err.Error())
		}
		products = append(products, product)
	}

	return &products, nil
}
