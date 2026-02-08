package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/services"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// HandleCheckout buat handle POST /api/checkout
func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Checkout godoc
// @Summary Proses checkout transaksi
// @Description Membuat transaksi baru dengan multiple items. Akan mengurangi stock produk dan menghitung total amount.
// @Tags transactions
// @Accept json
// @Produce json
// @Param checkout body models.CheckoutRequest true "Data checkout berisi items (product_id dan quantity)"
// @Success 200 {object} models.Transaction "Transaksi berhasil dibuat"
// @Failure 400 {string} string "Bad Request - items kosong atau product tidak ditemukan"
// @Failure 500 {string} string "Internal Server Error - stock tidak cukup"
// @Router /api/checkout [post]
func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(req.Items) == 0 {
		http.Error(w, "Items cannot be empty", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
