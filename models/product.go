package models

import "time"

// Product itu struct buat nyimpen data produk
type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

// Transaction itu struct buat nyimpen data transaksi
type Transaction struct {
	ID          int                 `json:"id"`
	TotalAmount int                 `json:"total_amount"`
	CreatedAt   time.Time           `json:"created_at"`
	Details     []TransactionDetail `json:"details"`
}

// TransactionDetail itu struct buat nyimpen detail transaksi
type TransactionDetail struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name,omitempty"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}

// CheckoutItem itu struct buat item yang akan di checkout
type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// CheckoutRequest itu struct buat request checkout
type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

// DailySalesReport itu struct buat laporan penjualan harian
type DailySalesReport struct {
	TotalRevenue    int         `json:"total_revenue"`
	TotalTransaksi  int         `json:"total_transaksi"`
	ProdukTerlaris  *TopProduct `json:"produk_terlaris"`
}

// TopProduct itu struct buat produk terlaris
type TopProduct struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}
