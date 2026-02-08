package repositories

import (
	"database/sql"
	"fmt"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
)

type TransactionRepository struct {
	db *sql.DB
}

// NewTransactionRepository buat bikin instance repository baru
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// CreateTransaction buat bikin transaksi baru dengan multiple items
func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		// Check stock
		if stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)", productName, stock, item.Quantity)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// FIX: Insert transaction details dengan RETURNING id untuk dapetin ID nya
	for i := range details {
		details[i].TransactionID = transactionID
		var detailID int
		err = tx.QueryRow(
			"INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal,
		).Scan(&detailID)
		if err != nil {
			return nil, err
		}
		details[i].ID = detailID
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
