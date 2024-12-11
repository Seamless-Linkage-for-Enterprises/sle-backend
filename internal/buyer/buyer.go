package buyer

import (
	"context"
	"sle/helpers"
	"time"
)

type Buyer struct {
	ID                string    `json:"s_id"`
	First_Name        string    `json:"s_first_name"`
	Last_Name         string    `json:"s_last_name"`
	Email             string    `json:"s_email"`
	Phone             string    `json:"s_phone"`
	Password          string    `json:"s_password"`
	Image_URL         string    `json:"s_image_url"`
	Address           string    `json:"s_address"`
	DOB               string    `json:"s_dob"`
	Is_Phone_Verified bool      `json:"is_phone_verified"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}

func NewBuyer(first_name, last_name, email, phone, password, image, address, dob string) *Buyer {
	current_time, _ := helpers.GetTime()

	return &Buyer{
		First_Name:        first_name,
		Last_Name:         last_name,
		Email:             email,
		Password:          password,
		Address:           address,
		Image_URL:         image,
		Phone:             phone,
		DOB:               dob,
		Is_Phone_Verified: false,
		Created_at:        current_time,
		Updated_at:        current_time,
	}
}

type CreateBuyerReq struct {
	First_Name string `json:"b_first_name"`
	Last_Name  string `json:"b_last_name"`
	Email      string `json:"b_email"`
	Password   string `json:"b_password"`
	Image_URL  string `json:"b_image_url"`
	Address    string `json:"b_address"`
	Phone      string `json:"b_phone"`
	DOB        string `json:"b_dob"`
}

type BuyerLoginReq struct {
	Phone    string `json:"b_phone"`
	Password string `json:"b_password"`
}

type BuyerForgetPasswordReq struct {
	Phone           string `json:"b_phone"`
	Password        string `json:"b_password"`
	ConfirmPassword string `json:"b_confirm_password"`
}

type Repository interface {
	BuyerSignup(ctx context.Context, buyer *Buyer) (*Buyer, error)
	GetBuyerByPhone(ctx context.Context, phone string) (*Buyer, error)
	GetBuyerByID(ctx context.Context, sid string) (*Buyer, error)
	GetAllBuyers(ctx context.Context, page int, recordPerPage int) (*[]Buyer, error)
	BuyerForgetPassword(ctx context.Context, phone, password string) error
	DeleteBuyer(ctx context.Context, id string) error
	IsBuyerVerified(ctx context.Context, phone string) error
	VerifyBuyer(ctx context.Context, phone string) error
}

type Service interface {
	BuyerSignup(ctx context.Context, req CreateBuyerReq) (*Buyer, error)
	BuyerLogin(ctx context.Context, req BuyerLoginReq) (*Buyer, error)
	GetBuyerByID(ctx context.Context, sid string) (*Buyer, error)
	GetBuyerByPhone(ctx context.Context, phone string) (*Buyer, error)
	GetAllBuyers(ctx context.Context, page int, recordPerPage int) (*[]Buyer, error)
	BuyerForgetPassword(ctx context.Context, req BuyerForgetPasswordReq) error
	DeleteBuyer(ctx context.Context, id string) error
	IsBuyerVerified(ctx context.Context, phone string) error
	VerifyBuyer(ctx context.Context, phone string) error
}
