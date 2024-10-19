package bookmark

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewBookmarkRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

// * create bookmark
func (r *repository) CreateBookmark(ctx context.Context, req Bookmark) error {
	// First, check if the bookmark already exists
	existsQuery := `
	SELECT EXISTS (
		SELECT 1 FROM bookmarks WHERE p_id = $1 AND b_id = $2
	)`
	var exists bool
	err := r.db.QueryRow(ctx, existsQuery, req.ProductID, req.BuyerID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("bookmark already added")
	}

	// If it doesn't exist, insert the new bookmark
	insertQuery := `
	INSERT INTO bookmarks (p_id, b_id) 
	VALUES ($1, $2)
	`
	_, err = r.db.Exec(ctx, insertQuery, req.ProductID, req.BuyerID)
	if err != nil {
		return err
	}

	return nil
}

// * delete product
func (r *repository) DeleteBookmark(ctx context.Context, bookmark_id string) error {
	query := `
	DELETE FROM bookmarks WHERE bookmark_id=$1
	`

	_, err := r.db.Exec(ctx, query, bookmark_id)
	if err != nil {
		return err
	}
	return nil
}

// * get all bookmarks by buyer
func (r *repository) GetAllBookmarks(ctx context.Context, buyerID string, page, recordPerPage int) (*[]BookmarkRes, error) {

	query := `
SELECT 
    b.bookmark_id,
        b.b_id,
        p.p_id,
        p.p_name,
        p.p_category,
        p.p_brand,
        p.p_status,
        p.p_description,
        p.p_quantity,
        p.p_image,
        p.p_price,
        p.created_at,
        p.updated_at,
        s.s_id,
        s.s_first_name,
        s.s_last_name,
        s.s_phone
FROM 
    bookmarks b
JOIN 
    products p ON b.p_id = p.p_id
JOIN 
    sellers s ON p.s_id = s.s_id
WHERE 
    b.b_id = $1;
    `

	rows, err := r.db.Query(ctx, query, buyerID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var bookmarks []BookmarkRes
	for rows.Next() {
		var bookmark BookmarkRes
		if err := rows.Scan(
			&bookmark.BookmarkID,
			&bookmark.BuyerID,
			&bookmark.ProductID,
			&bookmark.Name,
			&bookmark.Category,
			&bookmark.Brand,
			&bookmark.Status,
			&bookmark.Description,
			&bookmark.Quantity,
			&bookmark.Image,
			&bookmark.Price,
			&bookmark.Created_at,
			&bookmark.Updated_at,
			&bookmark.SID,
			&bookmark.S_First_Name,
			&bookmark.S_Last_Name,
			&bookmark.S_Phone,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return nil, err
	}

	log.Printf("Number of bookmarks fetched: %d", len(bookmarks))

	if len(bookmarks) == 0 {
		log.Println("No bookmarks found for the given buyer ID.")
		return &[]BookmarkRes{}, nil // Return an empty slice
	}

	return &bookmarks, nil
}
