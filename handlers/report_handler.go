package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/services"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

// HandleReportHariIni buat handle GET /api/report/hari-ini
func (h *ReportHandler) HandleReportHariIni(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetDailySales(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleReport buat handle GET /api/report dengan date range
func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetReportByDateRange(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetDailySales godoc
// @Summary Laporan penjualan hari ini
// @Description Mendapatkan ringkasan penjualan hari ini: total revenue, total transaksi, dan produk terlaris.
// @Tags reports
// @Accept json
// @Produce json
// @Success 200 {object} models.DailySalesReport "Laporan berisi total_revenue, total_transaksi, produk_terlaris"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/report/hari-ini [get]
func (h *ReportHandler) GetDailySales(w http.ResponseWriter, r *http.Request) {
	report, err := h.service.GetDailySales()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

// GetReportByDateRange godoc
// @Summary Laporan penjualan berdasarkan periode
// @Description Mendapatkan ringkasan penjualan dalam rentang tanggal tertentu. Optional challenge dari Bootcamp Session 3.
// @Tags reports
// @Accept json
// @Produce json
// @Param start_date query string true "Tanggal mulai (format: YYYY-MM-DD, contoh: 2026-01-01)"
// @Param end_date query string true "Tanggal akhir (format: YYYY-MM-DD, contoh: 2026-02-01)"
// @Success 200 {object} models.DailySalesReport "Laporan berisi total_revenue, total_transaksi, produk_terlaris"
// @Failure 400 {string} string "Bad Request - format tanggal salah"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/report [get]
func (h *ReportHandler) GetReportByDateRange(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	report, err := h.service.GetReportByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
