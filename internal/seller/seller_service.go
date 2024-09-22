package seller

import (
	"context"
	"errors"
	"log"
	"sle/helpers"
	"sle/service/email"
	"time"
)

type service struct {
	Repository
	emailRepo email.Repository
	timeout   time.Duration
}

func NewSellerService(r Repository, emailRepo email.Repository) Service {
	return &service{Repository: r, emailRepo: emailRepo, timeout: time.Duration(100) * time.Second}
}

func (s *service) SellerSignup(ctx context.Context, req CreateSellerReq) (*Seller, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	seller := Seller{
		First_Name:        req.First_Name,
		Last_Name:         req.Last_Name,
		Email:             req.Email,
		Password:          req.Password,
		Image_URL:         req.Image_URL,
		Address:           req.Address,
		Phone:             req.Phone,
		PAN_Card:          req.PAN_Card,
		DOB:               req.DOB,
		Description:       req.Description,
		GST_Number:        req.GST_Number,
		Company_Name:      req.Company_Name,
		Is_Verified:       true,
		Is_Email_Verified: true,
	}

	res, err := s.Repository.SellerSignup(ctx, &seller)
	if err != nil {
		return nil, err
	}

	// // generate otp
	// if err := s.emailRepo.Generateotp(res.First_Name, res.Last_Name, res.Email, res.ID); err != nil {
	// 	log.Println("Failed to send otp", err.Error())
	// }

	return res, nil
}

func (s *service) GetSellerByID(ctx context.Context, sid string) (*Seller, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	seller, err := s.Repository.GetSellerByID(ctx, sid)
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (s *service) GetAllSellers(ctx context.Context, page int, recordPerPage int) (*[]Seller, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	sellers, err := s.Repository.GetAllSellers(ctx, page, recordPerPage)
	if err != nil {
		return nil, err
	}

	return sellers, nil
}

func (s *service) SellerLogin(ctx context.Context, req SellerLoginReq) (*Seller, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	seller, err := s.Repository.GetSellerByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	_, err = helpers.VerifyPassword(req.Password, seller.Password)
	log.Println(req.Password, seller.Password)
	if err != nil {
		return nil, errors.New("password was wrong")
	}

	return seller, nil
}

func (s *service) SellerForgetPassword(ctx context.Context, req SellerForgetPasswordReq) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// hash new password
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return err
	}

	if err := s.Repository.SellerForgetPassword(ctx, req.Email, hashedPassword); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteSeller(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	return s.Repository.DeleteSeller(ctx, id)
}
