package seller

import (
	"errors"
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

func NewSellerHandler(s Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) SellerSignup(c *gin.Context) (int, error) {
	var req CreateSellerReq

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

	seller, err := h.Service.SellerSignup(c.Request.Context(), req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, seller)
}

func (h *Handler) GetSellerByID(c *gin.Context) (int, error) {
	s_id := c.Param("sid")

	seller, err := h.Service.GetSellerByID(c.Request.Context(), s_id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, seller)
}

func (h *Handler) GetAllSellers(c *gin.Context) (int, error) {
	page, errP := strconv.Atoi(c.Query("page"))
	recordPerPage, errO := strconv.Atoi(c.Query("recordPerPage"))

	// handling default page(offset) and recordPerPage(limit)
	if errP != nil || page < 0 {
		page = 0
	}

	if errO != nil || recordPerPage < 10 {
		recordPerPage = 10
	}

	seller, err := h.Service.GetAllSellers(c.Request.Context(), page, recordPerPage)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, seller)
}

func (h *Handler) SellerLogin(c *gin.Context) (int, error) {
	var req SellerLoginReq

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, err
	}

	msg, ok := validateLoginDetails(&req)
	if !ok {
		return http.StatusBadRequest, errors.New(msg)
	}

	seller, err := h.Service.SellerLogin(c.Request.Context(), req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// validate password
	if _, err = helpers.VerifyPassword(req.Password, seller.Password); err != nil {
		return http.StatusUnauthorized, err
	}

	return api.WriteData(c, seller)
}

func (h *Handler) SellerForgetPassword(c *gin.Context) (int, error) {
	var req SellerForgetPasswordReq

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
	if err := h.Service.SellerForgetPassword(c.Request.Context(), req); err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteMessage(c, "password was updated.")
}

func validateSignUpDetails(s *CreateSellerReq) (string, bool) {
	var errors []string

	// Validate required fields
	if s.First_Name == "" {
		errors = append(errors, "first_name required")
	}

	if s.Last_Name == "" {
		errors = append(errors, "last_name required")
	}

	if !utils.ValidateEmail(s.Email) {
		errors = append(errors, "provide valid email")
	}

	if str, ok := utils.ValidatePassword(s.Password); !ok {
		errors = append(errors, str)
	}

	if !utils.ValidateURL(s.Image_URL) {
		errors = append(errors, "image_url required")
	}

	if s.Address == "" {
		errors = append(errors, "address required")
	}

	if len(s.Phone) != 10 {
		errors = append(errors, "phone must be 10 digits")
	}

	// Validate PAN card
	if !utils.ValidatePANCard(s.PAN_Card) {
		errors = append(errors, "provide valid PAN card")
	}

	if s.DOB == "" {
		errors = append(errors, "dob required")
	}

	if s.Description == "" {
		errors = append(errors, "description required")
	}

	// If there are errors, join them into a single string
	if len(errors) > 0 {
		return "Validation failed: " + strings.Join(errors, " / "), false
	}

	return "", true
}

func validateLoginDetails(s *SellerLoginReq) (string, bool) {
	var errors []string

	if !utils.ValidateEmail(s.Email) {
		errors = append(errors, "provide valid email")
	}

	if len(errors) > 0 {
		return "Validation failed: " + strings.Join(errors, " / "), false
	}

	return "", true
}

func validateForgetPasswordDetails(s *SellerForgetPasswordReq) (string, bool) {

	var errors []string

	if !utils.ValidateEmail(s.Email) {
		errors = append(errors, "provide valid email")
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
