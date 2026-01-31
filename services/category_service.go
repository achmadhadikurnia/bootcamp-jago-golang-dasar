package services

import (
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

// NewCategoryService buat bikin instance service baru
func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetAll buat ambil semua categories
func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

// GetByID buat ambil category by ID
func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

// Create buat bikin category baru
func (s *CategoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

// Update buat update category
func (s *CategoryService) Update(category *models.Category) error {
	return s.repo.Update(category)
}

// Delete buat hapus category
func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
