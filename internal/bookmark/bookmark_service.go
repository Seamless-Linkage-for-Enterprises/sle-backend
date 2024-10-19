package bookmark

import (
	"context"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewBookmarkService(r Repository) Service {
	return &service{Repository: r, timeout: time.Duration(100) * time.Second}
}

func (s *service) CreateBookmark(ctx context.Context, req Bookmark) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.CreateBookmark(ctx, req)
}

func (s *service) DeleteBookmark(ctx context.Context, bookmark_id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.DeleteBookmark(ctx, bookmark_id)
}

func (s *service) GetAllBookmarks(ctx context.Context, buyer_id string, page, recordPerPage int) (*[]BookmarkRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.Repository.GetAllBookmarks(ctx, buyer_id, page, recordPerPage)
}
