package seller

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewSellerRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) SellerSignup(ctx context.Context, seller *Seller) (*Seller, error) {
	query := `INSERT INTO sellers (s_first_name,s_last_name,s_email,s_password,s_image_url,s_address,s_phone,s_pan_card,s_dob,s_company_name,s_description,s_gst_number) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) 
	RETURNING s_id,created_at,updated_at`

	err := r.db.QueryRow(ctx, query, seller.First_Name, seller.Last_Name, seller.Email, seller.Password, seller.Image_URL, seller.Address, seller.Phone, seller.PAN_Card, seller.DOB, seller.Company_Name, seller.Description, seller.GST_Number).Scan(&seller.ID, &seller.Created_at, &seller.Updated_at)
	if err != nil {
		return nil, err
	}

	if seller.ID == "" {
		return nil, errors.New("failed to insert record")
	}

	seller.Is_Verified = true // initialy it is false
	seller.Is_Email_Verified = true
	seller.Password = ""
	return seller, nil
}

func (r *repository) GetSellerByID(ctx context.Context, sid string) (*Seller, error) {
	query := `
	SELECT s_id,s_first_name,s_last_name,s_email,s_image_url,s_address,s_phone,s_pan_card,s_dob,s_company_name,s_description,is_verified,is_email_verified,s_gst_number,created_at,updated_at
	FROM sellers
	WHERE s_id = $1
	`
	var seller Seller
	var date time.Time

	if err := r.db.QueryRow(ctx, query, sid).Scan(
		&seller.ID,
		&seller.First_Name,
		&seller.Last_Name,
		&seller.Email,
		&seller.Image_URL,
		&seller.Address,
		&seller.Phone,
		&seller.PAN_Card,
		&date,
		&seller.Company_Name,
		&seller.Description,
		&seller.Is_Verified,
		&seller.Is_Email_Verified,
		&seller.GST_Number,
		&seller.Created_at,
		&seller.Updated_at); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("seller doesn't exist with the provided s_id")
		}
		return nil, err
	}

	seller.DOB = date.Format("2006-01-02")

	return &seller, nil
}

func (r *repository) GetAllSellers(ctx context.Context, page int, recordPerPage int) (*[]Seller, error) {
	query := `
	SELECT s_id,s_first_name,s_last_name,s_email,s_image_url,s_address,s_phone,s_pan_card,s_dob,s_company_name,s_description,is_verified,is_email_verified,s_gst_number,created_at,updated_at
	FROM sellers
	OFFSET $1 
	LIMIT $2
	`
	var sellers []Seller
	var date time.Time

	row, err := r.db.Query(ctx, query, page, recordPerPage)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var seller Seller
		if err := row.Scan(
			&seller.ID,
			&seller.First_Name,
			&seller.Last_Name,
			&seller.Email,
			&seller.Image_URL,
			&seller.Address,
			&seller.Phone,
			&seller.PAN_Card,
			&date,
			&seller.Company_Name,
			&seller.Description,
			&seller.Is_Verified,
			&seller.Is_Email_Verified,
			&seller.GST_Number,
			&seller.Created_at,
			&seller.Updated_at); err != nil {
			log.Println(err.Error())
		}
		seller.DOB = date.Format("2006-01-02")
		sellers = append(sellers, seller)
	}

	return &sellers, nil
}

func (r *repository) GetSellerByEmail(ctx context.Context, email string) (*Seller, error) {
	query := `
	SELECT s_id,s_first_name,s_last_name,s_email,s_password,s_image_url,s_address,s_phone,s_pan_card,s_dob,s_company_name,s_description,is_verified,is_email_verified,s_gst_number,created_at,updated_at
	FROM sellers
	WHERE s_email=$1`

	var seller Seller
	var date time.Time

	if err := r.db.QueryRow(ctx, query, email).Scan(
		&seller.ID,
		&seller.First_Name,
		&seller.Last_Name,
		&seller.Email,
		&seller.Password,
		&seller.Image_URL,
		&seller.Address,
		&seller.Phone,
		&seller.PAN_Card,
		&date,
		&seller.Company_Name,
		&seller.Description,
		&seller.Is_Verified,
		&seller.Is_Email_Verified,
		&seller.GST_Number,
		&seller.Created_at,
		&seller.Updated_at); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("seller doesn't exist with the provided s_email")
		}
		log.Println(err.Error())
		return nil, err
	}

	seller.DOB = date.Format("2006-01-02")

	return &seller, nil
}

func (r *repository) SellerForgetPassword(ctx context.Context, email, password string) error {

	// check seller exist with email
	_, err := r.GetSellerByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("seller doesn't exist with the provided s_email")
		}
		return err
	}

	query := `
	UPDATE sellers 
	SET s_password=$1
	WHERE s_email=$2
	`

	// update seller password
	if _, err := r.db.Exec(ctx, query, password, email); err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteSeller(ctx context.Context, id string) error {
	query := `
	DELETE FROM sellers WHERE s_id=$1
	`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
