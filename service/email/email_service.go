package email

type service struct {
	Repository
}

func NewEmailService(r Repository) Service {
	return &service{Repository: r}
}

func (s *service) Generateotp(first_name, last_name, email, s_id string) error {
	if err := s.Repository.Generateotp(first_name, last_name, email, s_id); err != nil {
		return err
	}

	return nil
}

func (s *service) VerifyOTP(otp, email string) error {
	if err := s.Repository.VerifyOTP(otp, email); err != nil {
		return err
	}

	// update is_email_verify = true
	if err := s.Repository.UpdateIsEmailUpdated(email); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteExpiredOTPs() error {
	return s.Repository.DeleteExpiredOTPs()
}
