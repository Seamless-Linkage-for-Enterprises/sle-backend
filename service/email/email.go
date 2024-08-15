package email

type OTP struct {
	SID        string `json:"s_id"`
	OTP        string `json:"otp"`
	Email      string `json:"email"`
	Expires_at int    `json:"expires_at"`
}

type VerifyOTPReq struct {
	OTP   string `json:"otp"`
	Email string `json:"email"`
}

type ResendOTPReq struct {
	Email      string `json:"email"`
	SID        string `json:"s_id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}

type Repository interface {
	Generateotp(first_name, last_name, email, s_id string) error
	VerifyOTP(otp, email string) error
	UpdateIsEmailUpdated(s_id string) error
	DeleteExpiredOTPs() error
}

type Service interface {
	Generateotp(first_name, last_name, email, s_id string) error
	VerifyOTP(otp, email string) error
	DeleteExpiredOTPs() error
}
