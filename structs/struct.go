package structs

import (
	"time"
)

type Book struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Image_url    string    `json:"image_url"`
	Release_year int       `json:"release_year"`
	Price        int       `json:"price"`
	Total_page   int       `json:"total_page"`
	Thickness    string    `json:"thickness"`
	Category_id  int       `json:"category_id"`
	Created_at   time.Time `json:"created_at"`
	Created_by   string    `json:"created_by"`
	Modified_at  time.Time `json:"modified_at"`
	Modified_by  string    `json:"modified_by"`
}

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Created_at  time.Time `json:"created_at"`
	Created_by  string    `json:"created_by"`
	Modified_at time.Time `json:"modified_at"`
	Modified_by string    `json:"modified_by"`
}

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Created_at  time.Time `json:"created_at"`
	Created_by  string    `json:"created_by"`
	Modified_at time.Time `json:"modified_at"`
	Modified_by string    `json:"modified_by"`
}
