package repositories

import (
	"database/sql"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
)

type ReportRepository struct {
	db *sql.DB
}

// NewReportRepository buat bikin instance repository baru
func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// GetDailySales buat ambil laporan penjualan hari ini
func (r *ReportRepository) GetDailySales() (*models.DailySalesReport, error) {
	report := &models.DailySalesReport{}

	// Query untuk total revenue dan total transaksi hari ini
	queryTotal := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`
	err := r.db.QueryRow(queryTotal).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// Query untuk produk terlaris hari ini
	queryTop := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`
	topProduct := &models.TopProduct{}
	err = r.db.QueryRow(queryTop).Scan(&topProduct.Nama, &topProduct.QtyTerjual)
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = nil
	} else if err != nil {
		return nil, err
	} else {
		report.ProdukTerlaris = topProduct
	}

	return report, nil
}

// GetReportByDateRange buat ambil laporan penjualan berdasarkan range tanggal
func (r *ReportRepository) GetReportByDateRange(startDate, endDate string) (*models.DailySalesReport, error) {
	report := &models.DailySalesReport{}

	// Query untuk total revenue dan total transaksi dalam range
	queryTotal := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE DATE(created_at) >= $1 AND DATE(created_at) <= $2
	`
	err := r.db.QueryRow(queryTotal, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// Query untuk produk terlaris dalam range
	queryTop := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) >= $1 AND DATE(t.created_at) <= $2
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`
	topProduct := &models.TopProduct{}
	err = r.db.QueryRow(queryTop, startDate, endDate).Scan(&topProduct.Nama, &topProduct.QtyTerjual)
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = nil
	} else if err != nil {
		return nil, err
	} else {
		report.ProdukTerlaris = topProduct
	}

	return report, nil
}
