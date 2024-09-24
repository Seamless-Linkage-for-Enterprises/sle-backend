package buyer

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

func NewBuyerRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) BuyerSignup(ctx context.Context, buyer *Buyer) (*Buyer, error) {
	query := `INSERT INTO buyers (b_first_name,b_last_name,b_email,b_password,b_image_url,b_address,b_phone,b_dob) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8) 
	RETURNING b_id,created_at,updated_at`

	err := r.db.QueryRow(ctx, query, buyer.First_Name, buyer.Last_Name, buyer.Email, buyer.Password, buyer.Image_URL, buyer.Address, buyer.Phone, buyer.DOB).Scan(&buyer.ID, &buyer.Created_at, &buyer.Updated_at)
	if err != nil {
		return nil, err
	}

	if buyer.ID == "" {
		return nil, errors.New("failed to insert record")
	}

	buyer.Is_Phone_Verified = false
	buyer.Password = ""
	return buyer, nil
}

func (r *repository) GetBuyerByID(ctx context.Context, sid string) (*Buyer, error) {
	query := `
	SELECT b_id,b_first_name,b_last_name,b_email,b_password,b_image_url,b_address,b_phone,b_dob,is_phone_verified,created_at,updated_at
	FROM buyers
	WHERE b_id = $1
	`
	var buyer Buyer
	var date time.Time

	if err := r.db.QueryRow(ctx, query, sid).Scan(
		&buyer.ID,
		&buyer.First_Name,
		&buyer.Last_Name,
		&buyer.Email,
		&buyer.Password,
		&buyer.Image_URL,
		&buyer.Address,
		&buyer.Phone,
		&date,
		&buyer.Is_Phone_Verified,
		&buyer.Created_at,
		&buyer.Updated_at); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("buyer doesn't exist with the provided b_id")
		}
		return nil, err
	}

	buyer.DOB = date.Format("2006-01-02")

	return &buyer, nil
}

func (r *repository) GetAllBuyers(ctx context.Context, page int, recordPerPage int) (*[]Buyer, error) {
	query := `
	SELECT b_id,b_first_name,b_last_name,b_email,b_image_url,b_address,b_phone,b_dob,is_phone_verified,created_at,updated_at
	FROM buyers
	OFFSET $1 
	LIMIT $2
	`
	var buyers []Buyer
	var date time.Time

	row, err := r.db.Query(ctx, query, page, recordPerPage)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var buyer Buyer
		if err := row.Scan(
			&buyer.ID,
			&buyer.First_Name,
			&buyer.Last_Name,
			&buyer.Email,
			&buyer.Image_URL,
			&buyer.Address,
			&buyer.Phone,
			&date,
			&buyer.Is_Phone_Verified,
			&buyer.Created_at,
			&buyer.Updated_at); err != nil {
			log.Println(err.Error())
		}
		buyer.DOB = date.Format("2006-01-02")
		buyers = append(buyers, buyer)
	}

	return &buyers, nil
}

func (r *repository) GetBuyerByPhone(ctx context.Context, email string) (*Buyer, error) {
	query := `
	SELECT b_id,b_first_name,b_last_name,b_email,b_password,b_image_url,b_address,b_phone,b_dob,is_phone_verified,created_at,updated_at 
	FROM buyers
	WHERE b_phone=$1`

	var buyer Buyer
	var date time.Time

	if err := r.db.QueryRow(ctx, query, email).Scan(
		&buyer.ID,
		&buyer.First_Name,
		&buyer.Last_Name,
		&buyer.Email,
		&buyer.Password,
		&buyer.Image_URL,
		&buyer.Address,
		&buyer.Phone,
		&date,
		&buyer.Is_Phone_Verified,
		&buyer.Created_at,
		&buyer.Updated_at); err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("buyer doesn't exist with the provided b_phone")
		}
		log.Println(err.Error())
		return nil, err
	}

	buyer.DOB = date.Format("2006-01-02")

	return &buyer, nil
}

func (r *repository) BuyerForgetPassword(ctx context.Context, phone, password string) error {

	// check buyer exist with email
	_, err := r.GetBuyerByPhone(ctx, phone)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("buyer doesn't exist with the provided b_phone")
		}
		return err
	}

	query := `
	UPDATE buyers 
	SET b_password=$1
	WHERE b_phone=$2
	`

	// update buyer password
	if _, err := r.db.Exec(ctx, query, password, phone); err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteBuyer(ctx context.Context, id string) error {
	query := `
	DELETE FROM buyers WHERE b_id=$1
	`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) IsBuyerVerified(ctx context.Context, id string) error {
	query := `
	SELECT b_id,b_first_name,b_last_name,b_email,b_image_url,b_address,b_phone,b_dob,is_phone_verified,created_at,updated_at 
	FROM buyers
	WHERE b_id=$1`

	var buyer Buyer
	var date time.Time

	if err := r.db.QueryRow(ctx, query, id).Scan(
		&buyer.ID,
		&buyer.First_Name,
		&buyer.Last_Name,
		&buyer.Email,
		&buyer.Image_URL,
		&buyer.Address,
		&buyer.Phone,
		&date,
		&buyer.Is_Phone_Verified,
		&buyer.Created_at,
		&buyer.Updated_at); err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("buyer doesn't exist with the provided b_id")
		}
		log.Println(err.Error())
		return err
	}

	if !buyer.Is_Phone_Verified {
		return errors.New("buyer doesn't verified")
	}

	return nil
}

func (r *repository) VerifyBuyer(ctx context.Context, phone string) error {

	query := `
	UPDATE buyers 
	SET is_phone_verified=$1
	WHERE b_phone=$2
	`

	// update buyer verified
	pg, err := r.db.Exec(ctx, query, true, phone)
	if err != nil {
		return err
	}

	if pg.RowsAffected() <= 0 {
		return errors.New("buyer doesn't exists with phone")
	}

	return nil
}
