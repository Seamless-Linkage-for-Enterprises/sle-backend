package buyer

import (
	"context"
	"errors"
	"log"
	"sle/helpers"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewBuyerService(r Repository) Service {
	return &service{Repository: r, timeout: time.Duration(100) * time.Second}
}

func (s *service) BuyerSignup(ctx context.Context, req CreateBuyerReq) (*Buyer, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	buyer := Buyer{
		First_Name:        req.First_Name,
		Last_Name:         req.Last_Name,
		Email:             req.Email,
		Password:          req.Password,
		Image_URL:         req.Image_URL,
		Address:           req.Address,
		Phone:             req.Phone,
		DOB:               req.DOB,
		Is_Phone_Verified: true,
	}

	res, err := s.Repository.BuyerSignup(ctx, &buyer)
	if err != nil {
		return nil, err
	}

	// // generate otp
	// if err := s.emailRepo.Generateotp(res.First_Name, res.Last_Name, res.Email, res.ID); err != nil {
	// 	log.Println("Failed to send otp", err.Error())
	// }

	return res, nil
}

func (s *service) GetBuyerByID(ctx context.Context, bid string) (*Buyer, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	buyer, err := s.Repository.GetBuyerByID(ctx, bid)
	if err != nil {
		return nil, err
	}

	return buyer, nil
}

func (s *service) GetAllBuyers(ctx context.Context, page int, recordPerPage int) (*[]Buyer, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	buyers, err := s.Repository.GetAllBuyers(ctx, page, recordPerPage)
	if err != nil {
		return nil, err
	}

	return buyers, nil
}

func (s *service) BuyerLogin(ctx context.Context, req BuyerLoginReq) (*Buyer, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	buyer, err := s.Repository.GetBuyerByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}

	_, err = helpers.VerifyPassword(req.Password, buyer.Password)
	log.Println(req.Password, buyer.Password)
	if err != nil {
		return nil, errors.New("password was wrong")
	}

	return buyer, nil
}

func (s *service) BuyerForgetPassword(ctx context.Context, req BuyerForgetPasswordReq) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// hash new password
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return err
	}

	if err := s.Repository.BuyerForgetPassword(ctx, req.Phone, hashedPassword); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteBuyer(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.DeleteBuyer(ctx, id)
}

func (s *service) IsBuyerVerified(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.IsBuyerVerified(ctx, id)
}

func (s *service) VerifyBuyer(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.VerifyBuyer(ctx, id)
}
