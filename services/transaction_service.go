package services

import (
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

// NewTransactionService buat bikin instance service baru
func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

// Checkout buat proses checkout items
func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}
