package repositories

import (
	"database/sql"
	"errors"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
)

type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository buat bikin instance repository baru
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetAll buat ambil semua products dari database
func (r *ProductRepository) GetAll() ([]models.Product, error) {
	query := `SELECT p.id, p.name, p.price, p.stock, p.category_id, COALESCE(c.name, '') as category_name
			  FROM products p
			  LEFT JOIN categories c ON p.category_id = c.id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.CategoryName)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// GetByID buat ambil product berdasarkan ID
func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	query := `SELECT p.id, p.name, p.price, p.stock, p.category_id, COALESCE(c.name, '') as category_name
			  FROM products p
			  LEFT JOIN categories c ON p.category_id = c.id
			  WHERE p.id = $1`

	var p models.Product
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.CategoryName)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Create buat bikin product baru
func (r *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

// Update buat update product yang udah ada
func (r *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := r.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

// Delete buat hapus product
func (r *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}
