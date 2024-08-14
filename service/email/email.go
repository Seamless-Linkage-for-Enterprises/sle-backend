package email

type OTP struct {
	SID        string
	OTP        string
	Email      string
	Expires_at int
}

type VerifyOTPReq struct {
	OTP   string
	Email string
}

type Repository interface {
	Generateotp(first_name, last_name, email, s_id string) error
	VerifyOTP(otp, email string) error
}

type Service interface {
	Generateotp(first_name, last_name, email, s_id string) error
	VerifyOTP(otp, email string) error
}
