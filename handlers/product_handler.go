package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/services"
)

type ProductHandler struct {
	service *services.ProductService
}

// NewProductHandler buat bikin instance handler baru
func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// HandleProducts buat handle GET /api/products dan POST /api/products
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetAll godoc
// @Summary Get all products
// @Description Get all products from database
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/products [get]
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Create godoc
// @Summary Create a new product
// @Description Create a new product in database
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product data"
// @Success 201 {object} models.Product
// @Failure 400 {string} string "Bad Request"
// @Router /api/products [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// HandleProductByID buat handle GET/PUT/DELETE /api/products/{id}
func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetByID godoc
// @Summary Get product by ID
// @Description Get a single product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Product not found"
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Update godoc
// @Summary Update a product
// @Description Update product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product data"
// @Success 200 {object} models.Product
// @Failure 400 {string} string "Bad Request"
// @Router /api/products/{id} [put]
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Delete godoc
// @Summary Delete a product
// @Description Delete product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/products/{id} [delete]
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product deleted successfully",
	})
}
