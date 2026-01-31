package models

// Category itu struct buat nyimpen data kategori
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
