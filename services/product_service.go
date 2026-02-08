package services

import (
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

// NewProductService buat bikin instance service baru
func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetAll buat ambil semua products
func (s *ProductService) GetAll(name string) ([]models.Product, error) {
	return s.repo.GetAll(name)
}

// GetByID buat ambil product by ID
func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

// Create buat bikin product baru
func (s *ProductService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

// Update buat update product
func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

// Delete buat hapus product
func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
