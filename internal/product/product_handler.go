package product

import (
	"net/http"
	api "sle/internal"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewProductHandler(s Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) CreateProduct(c *gin.Context) (int, error) {
	var req ProductReq

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, err
	}

	// validation

	product, err := h.Service.CreateProduct(c.Request.Context(), &req, req.SID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, product)

}

func (h *Handler) GetProductByID(c *gin.Context) (int, error) {
	pid := c.Param("pid")

	if pid == "" {
		return api.WriteMessage(c, "please provide pid")
	}

	products, err := h.Service.GetProductByID(c.Request.Context(), pid)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, products)
}

func (h *Handler) GetAllProducts(c *gin.Context) (int, error) {
	page, errP := strconv.Atoi(c.Query("page"))
	recordPerPage, errO := strconv.Atoi(c.Query("recordPerPage"))

	// handling default page(offset) and recordPerPage(limit)
	if errP != nil || page < 1 {
		page = 1
	}

	if errO != nil || recordPerPage < 10 {
		recordPerPage = 10
	}

	products, err := h.Service.GetAllProducts(c.Request.Context(), page, recordPerPage)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, products)
}

func (h *Handler) GetAllProductsBySellerAndCategory(c *gin.Context) (int, error) {
	page, errP := strconv.Atoi(c.Query("page"))
	recordPerPage, errO := strconv.Atoi(c.Query("recordPerPage"))
	s_id := c.Query("s_id")
	category := c.Query("category")

	// handling default page(offset) and recordPerPage(limit)
	if errP != nil || page < 1 {
		page = 1
	}

	if errO != nil || recordPerPage < 10 {
		recordPerPage = 10
	}

	products, err := h.Service.GetAllProductsBySellerAndCategory(c.Request.Context(), s_id, category, page, recordPerPage)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, products)
}

func (h *Handler) DeleteProduct(c *gin.Context) (int, error) {
	pid := c.Param("pid")

	if err := h.Service.DeleteProduct(c.Request.Context(), pid); err != nil {
		return http.StatusNotFound, err
	}

	return api.WriteMessage(c, "Product deleted.")
}

func (h *Handler) UpdateProductDetails(c *gin.Context) (int, error) {
	pid := c.Param("pid")

	var req ProductReq
	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, err
	}

	if err := h.Service.UpdateProductDetails(c.Request.Context(), pid, req); err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteMessage(c, "Product updated.")
}

func (h *Handler) UpdateStatus(c *gin.Context) (int, error) {
	pid := c.Param("pid")

	if err := h.Service.UpdateStatus(c.Request.Context(), pid); err != nil {
		return http.StatusNotFound, err
	}

	return api.WriteMessage(c, "Product status updated.")
}

func (h *Handler) SearchProduct(c *gin.Context) (int, error) {
	str := c.Param("str")
	page, errP := strconv.Atoi(c.Query("page"))
	recordPerPage, errO := strconv.Atoi(c.Query("recordPerPage"))

	// handling default page(offset) and recordPerPage(limit)
	if errP != nil || page < 1 {
		page = 1
	}

	if errO != nil || recordPerPage < 10 {
		recordPerPage = 10
	}

	products, err := h.Service.SearchProduct(c.Request.Context(), str, page, recordPerPage)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteData(c, products)
}
