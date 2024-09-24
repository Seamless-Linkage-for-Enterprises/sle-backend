package seller

import (
	"context"
	"sle/helpers"
	"time"
)

type Seller struct {
	ID                string    `json:"s_id"`
	First_Name        string    `json:"s_first_name"`
	Last_Name         string    `json:"s_last_name"`
	Email             string    `json:"s_email"`
	Password          string    `json:"s_password"`
	Image_URL         string    `json:"s_image_url"`
	Address           string    `json:"s_address"`
	Phone             string    `json:"s_phone"`
	PAN_Card          string    `json:"s_pan_card"`
	DOB               string    `json:"s_dob"`
	Company_Name      string    `json:"s_company_name"`
	Description       string    `json:"s_discription"`
	Is_Verified       bool      `json:"is_verified"`
	Is_Email_Verified bool      `json:"is_email_verified"`
	GST_Number        string    `json:"s_gst_number"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}

func NewSeller(first_name, last_name, email, password, image, address, pan_card, company_name, description, gst_number, phone, dob string) *Seller {
	current_time, _ := helpers.GetTime()
	return &Seller{
		First_Name:        first_name,
		Last_Name:         last_name,
		Email:             email,
		Password:          password,
		Address:           address,
		Image_URL:         image,
		Phone:             phone,
		PAN_Card:          pan_card,
		DOB:               dob,
		Company_Name:      company_name,
		Description:       description,
		Is_Verified:       true,
		Is_Email_Verified: true,
		GST_Number:        gst_number,
		Created_at:        current_time,
		Updated_at:        current_time,
	}
}

type CreateSellerReq struct {
	First_Name   string `json:"s_first_name"`
	Last_Name    string `json:"s_last_name"`
	Email        string `json:"s_email"`
	Password     string `json:"s_password"`
	Image_URL    string `json:"s_image_url"`
	Address      string `json:"s_address"`
	Phone        string `json:"s_phone"`
	PAN_Card     string `json:"s_pan_card"`
	DOB          string `json:"s_dob"`
	Company_Name string `json:"s_company_name"`
	Description  string `json:"s_description"`
	GST_Number   string `json:"s_gst_number"`
}

type SellerLoginReq struct {
	Email    string `json:"s_email"`
	Password string `json:"s_password"`
}

type SellerForgetPasswordReq struct {
	Email           string `json:"s_email"`
	Password        string `json:"s_password"`
	ConfirmPassword string `json:"s_confirm_password"`
}

type Repository interface {
	SellerSignup(ctx context.Context, seller *Seller) (*Seller, error)
	GetSellerByEmail(ctx context.Context, email string) (*Seller, error)
	GetSellerByID(ctx context.Context, sid string) (*Seller, error)
	GetAllSellers(ctx context.Context, page int, recordPerPage int) (*[]Seller, error)
	SellerForgetPassword(ctx context.Context, email, password string) error
	DeleteSeller(ctx context.Context, id string) error
}

type Service interface {
	SellerSignup(ctx context.Context, req CreateSellerReq) (*Seller, error)
	SellerLogin(ctx context.Context, req SellerLoginReq) (*Seller, error)
	GetSellerByID(ctx context.Context, sid string) (*Seller, error)
	GetAllSellers(ctx context.Context, page int, recordPerPage int) (*[]Seller, error)
	SellerForgetPassword(ctx context.Context, req SellerForgetPasswordReq) error
	DeleteSeller(id string) error
}
