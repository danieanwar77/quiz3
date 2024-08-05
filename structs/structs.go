package structs

import (
	"github.com/lib/pq"
)

type Book struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	CategoryId  int         `json:"category_id"`
	Description string      `json:"description"`
	ImageUrl    string      `json:"image_url"`
	ReleaseYear int         `json:"release_year"`
	Price       int         `json:"price"`
	TotalPage   int         `json:"total_page"`
	Thickness   string      `json:"thickness"`
	CreatedAt   pq.NullTime `json:"created_at"`
	CreatedBy   string      `json:"created_by"`
	ModifiedAt  pq.NullTime `json:"modified_at"`
	ModifiedBy  pq.NullTime `json:"modified_by"`
}

type Category struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	CreatedAt  pq.NullTime `json:"created_at"`
	CreatedBy  string      `json:"created_by"`
	ModifiedAt pq.NullTime `json:"modified_at"`
	ModifiedBy pq.NullTime `json:"modified_by"`
}

type Users struct {
	ID         int         `json:"id"`
	UserName   string      `json:"username"`
	Password   string      `json:"password"`
	CreatedAt  pq.NullTime `json:"created_at"`
	CreatedBy  string      `json:"created_by"`
	ModifiedAt pq.NullTime `json:"modified_at"`
	ModifiedBy pq.NullTime `json:"modified_by"`
}
