package services

import (
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

// NewReportService buat bikin instance service baru
func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

// GetDailySales buat ambil laporan penjualan hari ini
func (s *ReportService) GetDailySales() (*models.DailySalesReport, error) {
	return s.repo.GetDailySales()
}

// GetReportByDateRange buat ambil laporan penjualan berdasarkan range tanggal
func (s *ReportService) GetReportByDateRange(startDate, endDate string) (*models.DailySalesReport, error) {
	return s.repo.GetReportByDateRange(startDate, endDate)
}
