package email

import (
	"net/http"
	api "sle/internal"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewEmailHandler(s Service) Handler {
	return Handler{Service: s}
}

func (h *Handler) ResendOTP(c *gin.Context) (int, error) {
	var req ResendOTPReq

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, err
	}

	if err := h.Service.Generateotp(req.First_Name, req.Last_Name, req.Email, req.SID); err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteMessage(c, "email was sent")
}

func (h *Handler) VerifyOTP(c *gin.Context) (int, error) {
	var req VerifyOTPReq

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, err
	}

	if err := h.Service.VerifyOTP(req.OTP, req.Email); err != nil {
		return http.StatusUnauthorized, err
	}

	return api.WriteMessage(c, "your email is verified")
}

func (h *Handler) DeleteOTPs() error {
	return h.Service.DeleteExpiredOTPs()
}
