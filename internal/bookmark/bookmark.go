package bookmark

import (
	"context"
	"time"
)

type Bookmark struct {
	ProductID string `json:"p_id"`
	BuyerID   string `json:"b_id"`
}

type BookmarkRes struct {
	BookmarkID   string    `json:"bookmark_id"`
	BuyerID      string    `json:"b_id"`
	ProductID    string    `json:"p_id"`
	Name         string    `json:"p_name"`
	Category     string    `json:"p_category"`
	Brand        string    `json:"p_brand"`
	Status       bool      `json:"p_status"`
	Description  string    `json:"p_description"`
	Quantity     int       `json:"p_quantity"`
	Image        string    `json:"p_image"`
	Price        int       `json:"p_price"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	SID          string    `json:"s_id"`
	S_First_Name string    `json:"s_first_name"`
	S_Last_Name  string    `json:"s_last_name"`
	S_Phone      string    `json:"s_phone"`
}

type Repository interface {
	CreateBookmark(ctx context.Context, req Bookmark) error
	GetAllBookmarks(ctx context.Context, buyer_id string, page, recordPerPage int) (*[]BookmarkRes, error)
	DeleteBookmark(ctx context.Context, bookmark_id string) error
}

type Service interface {
	CreateBookmark(ctx context.Context, req Bookmark) error
	GetAllBookmarks(ctx context.Context, buyer_id string, page, recordPerPage int) (*[]BookmarkRes, error)
	DeleteBookmark(ctx context.Context, bookmark_id string) error
}
