package bookmark

import (
	"errors"
	"net/http"
	api "sle/internal"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewBookmarkHandler(s Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) CreateBookmark(c *gin.Context) (int, error) {

	pid := c.Param("pid") // product id
	bid := c.Param("bid") // buyer id

	if pid == "" || bid == "" {
		return http.StatusBadRequest, errors.New("product and buyer ids are required")
	}

	req := Bookmark{
		ProductID: pid,
		BuyerID:   bid,
	}

	if err := h.Service.CreateBookmark(c.Request.Context(), req); err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteMessage(c, "bookmark added")
}

func (h *Handler) DeleteBookmark(c *gin.Context) (int, error) {

	bookmark_id := c.Param("bookmarkid")

	if bookmark_id == "" {
		return http.StatusBadRequest, errors.New("bookmardid is required")
	}

	if err := h.Service.DeleteBookmark(c.Request.Context(), bookmark_id); err != nil {
		return http.StatusInternalServerError, err
	}

	return api.WriteMessage(c, "bookmark removed")
}

// * get all bookmark products
func (h *Handler) GetAllBookmarks(c *gin.Context) (int, error) {
	buyer_id := c.Param("bid")

	if buyer_id == "" {
		return http.StatusBadRequest, errors.New("buyer id is required")
	}
	page, errP := strconv.Atoi(c.Query("page"))
	recordPerPage, errO := strconv.Atoi(c.Query("recordPerPage"))

	// handling default page(offset) and recordPerPage(limit)
	if errP != nil || page < 1 {
		page = 1
	}

	if errO != nil || recordPerPage < 10 {
		recordPerPage = 10
	}

	products, err := h.Service.GetAllBookmarks(c.Request.Context(), buyer_id, page, recordPerPage)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return api.WriteData(c, products)
}
