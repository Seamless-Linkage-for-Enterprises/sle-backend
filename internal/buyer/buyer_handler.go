package buyer

import (
	"errors"
	"log"
	"net/http"
	"sle/helpers"
	api "sle/internal"
	"sle/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewBuyerHandler(s Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) BuyerSignup(c *gin.Context) (int, error) {
	var req CreateBuyerReq

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, err
	}

	msg, ok := validateSignUpDetails(&req)
	if !ok {
		return http.StatusBadRequest, errors.New(msg)
	}

	// TODO : check email domain means email exist in the world or not

	// hash password
	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return http.StatusInternalServerError, nil
	}

	req.Password = hashPassword

	seller, err := h.Service.BuyerSignup(c.Request.Context(), req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, seller)
}

func (h *Handler) GetBuyerByID(c *gin.Context) (int, error) {
	b_id := c.Param("bid")

	buyer, err := h.Service.GetBuyerByID(c.Request.Context(), b_id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, buyer)
}

func (h *Handler) GetAllBuyers(c *gin.Context) (int, error) {
	page, errP := strconv.Atoi(c.Query("page"))
	recordPerPage, errO := strconv.Atoi(c.Query("recordPerPage"))

	// handling default page(offset) and recordPerPage(limit)
	if errP != nil || page < 0 {
		page = 0
	}

	if errO != nil || recordPerPage < 10 {
		recordPerPage = 10
	}

	buyer, err := h.Service.GetAllBuyers(c.Request.Context(), page, recordPerPage)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, buyer)
}

func (h *Handler) BuyerLogin(c *gin.Context) (int, error) {
	var req BuyerLoginReq

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, err
	}

	msg, ok := validateLoginDetails(&req)
	if !ok {
		return http.StatusBadRequest, errors.New(msg)
	}

	seller, err := h.Service.BuyerLogin(c.Request.Context(), req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	log.Println(req.Password, seller.Password)

	// validate password
	if _, err = helpers.VerifyPassword(req.Password, seller.Password); err != nil {
		return http.StatusUnauthorized, err
	}

	return api.WriteData(c, seller)
}

func (h *Handler) BuyerForgetPassword(c *gin.Context) (int, error) {
	var req BuyerForgetPasswordReq

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, err
	}

	msg, ok := validateForgetPasswordDetails(&req)
	if !ok {
		return http.StatusBadRequest, errors.New(msg)
	}

	// compare password and hash password
	if req.Password != req.ConfirmPassword {
		return http.StatusBadRequest, errors.New("s_password and s_confirm_password must be same")
	}

	// update password
	if err := h.Service.BuyerForgetPassword(c.Request.Context(), req); err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteMessage(c, "password was updated.")
}

func (h *Handler) DeleteBuyer(c *gin.Context) (int, error) {
	id := c.Param("bid")

	if err := h.Service.DeleteBuyer(c.Request.Context(), id); err != nil {
		return http.StatusNotFound, err
	}

	return api.WriteMessage(c, "buyer deleted.")
}

func (h *Handler) IsBuyerVerified(c *gin.Context) (int, error) {
	id := c.Param("bid")

	if err := h.Service.IsBuyerVerified(c.Request.Context(), id); err != nil {
		return http.StatusNotFound, err
	}

	return api.WriteData(c, true)
}

func (h *Handler) VerifyBuyer(c *gin.Context) (int, error) {
	id := c.Param("bid")

	if err := h.Service.VerifyBuyer(c.Request.Context(), id); err != nil {
		return http.StatusNotFound, err
	}

	return api.WriteMessage(c, "buyer is verified.")
}

func validateSignUpDetails(b *CreateBuyerReq) (string, bool) {
	var errors []string

	// Validate required fields
	if b.First_Name == "" {
		errors = append(errors, "first_name required")
	}

	if b.Last_Name == "" {
		errors = append(errors, "last_name required")
	}

	if !utils.ValidateEmail(b.Email) {
		errors = append(errors, "provide valid email")
	}

	if str, ok := utils.ValidatePassword(b.Password); !ok {
		errors = append(errors, str)
	}

	if b.Image_URL == "" {
		errors = append(errors, "image_url required")
	}

	if b.Address == "" {
		errors = append(errors, "address required")
	}

	if len(b.Phone) != 10 {
		errors = append(errors, "phone must be 10 digits")
	}

	if b.DOB == "" {
		errors = append(errors, "dob required")
	}

	// If there are errors, join them into a single string
	if len(errors) > 0 {
		return "Validation failed: " + strings.Join(errors, " / "), false
	}

	return "", true
}

func validateLoginDetails(s *BuyerLoginReq) (string, bool) {
	var errors []string

	phone, _ := strconv.Atoi(s.Phone)

	if utils.CheckLength(phone, 10) {
		errors = append(errors, "provide 10 digits of phone number")
	}

	if len(errors) > 0 {
		return "Validation failed: " + strings.Join(errors, " / "), false
	}

	return "", true
}

func validateForgetPasswordDetails(s *BuyerForgetPasswordReq) (string, bool) {

	var errors []string

	phone, _ := strconv.Atoi(s.Phone)

	if utils.CheckLength(phone, 10) {
		errors = append(errors, "provide 10 digit phone number")
	}

	if msg, ok := utils.ValidatePassword(s.Password); !ok {
		errors = append(errors, msg)
	}

	if s.ConfirmPassword != s.Password {
		errors = append(errors, "confirm password and password must be same")
	}

	if len(errors) > 0 {
		return "Validation failed: " + strings.Join(errors, " / "), false
	}

	return "", true
}
