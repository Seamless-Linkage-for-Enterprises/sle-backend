package email

import "github.com/gin-gonic/gin"

type Handler struct {
	Service
}

func NewEmailHandler(s Service) Handler {
	return Handler{Service: s}
}

func (h *Handler) GenerateOTP(c *gin.Context) (int, error) {
	return 200, nil
}

func (h *Handler) VerifyOTP(c *gin.Context) (int, error) {
	return 200, nil
}
