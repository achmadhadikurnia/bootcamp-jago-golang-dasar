package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/models"
)

// CategoryHandler handles all category-related HTTP requests
type CategoryHandler struct {
	categories []models.Category
	nextID     int
	mu         sync.RWMutex
}

// NewCategoryHandler creates a new CategoryHandler instance with dummy data
func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		categories: []models.Category{
			{ID: 1, Name: "Electronics", Description: "Produk elektronik dan gadget"},
			{ID: 2, Name: "Fashion", Description: "Pakaian dan aksesoris"},
			{ID: 3, Name: "Books", Description: "Buku dan majalah"},
		},
		nextID: 4,
	}
}

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// CreateCategoryRequest represents the request body for creating/updating a category
type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// writeJSON writes a JSON response with the given status code
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// GetAllCategories handles GET /categories
func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	writeJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "Categories retrieved successfully",
		Data:    h.categories,
	})
}

// GetCategory handles GET /categories/{id}
func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request, id int) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, category := range h.categories {
		if category.ID == id {
			writeJSON(w, http.StatusOK, Response{
				Success: true,
				Message: "Category retrieved successfully",
				Data:    category,
			})
			return
		}
	}

	writeJSON(w, http.StatusNotFound, Response{
		Success: false,
		Message: "Category not found",
	})
}

// CreateCategory handles POST /categories
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	// Validate required fields
	if strings.TrimSpace(req.Name) == "" {
		writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Name is required",
		})
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	category := models.Category{
		ID:          h.nextID,
		Name:        req.Name,
		Description: req.Description,
	}
	h.nextID++
	h.categories = append(h.categories, category)

	writeJSON(w, http.StatusCreated, Response{
		Success: true,
		Message: "Category created successfully",
		Data:    category,
	})
}

// UpdateCategory handles PUT /categories/{id}
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request, id int) {
	var req CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	// Validate required fields
	if strings.TrimSpace(req.Name) == "" {
		writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Name is required",
		})
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	for i, category := range h.categories {
		if category.ID == id {
			h.categories[i].Name = req.Name
			h.categories[i].Description = req.Description

			writeJSON(w, http.StatusOK, Response{
				Success: true,
				Message: "Category updated successfully",
				Data:    h.categories[i],
			})
			return
		}
	}

	writeJSON(w, http.StatusNotFound, Response{
		Success: false,
		Message: "Category not found",
	})
}

// DeleteCategory handles DELETE /categories/{id}
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request, id int) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for i, category := range h.categories {
		if category.ID == id {
			// Remove the category from slice
			h.categories = append(h.categories[:i], h.categories[i+1:]...)

			writeJSON(w, http.StatusOK, Response{
				Success: true,
				Message: "Category deleted successfully",
			})
			return
		}
	}

	writeJSON(w, http.StatusNotFound, Response{
		Success: false,
		Message: "Category not found",
	})
}

// ServeHTTP implements http.Handler interface for routing
func (h *CategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/categories")
	path = strings.TrimSuffix(path, "/")

	// Handle /categories
	if path == "" {
		switch r.Method {
		case http.MethodGet:
			h.GetAllCategories(w, r)
		case http.MethodPost:
			h.CreateCategory(w, r)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, Response{
				Success: false,
				Message: "Method not allowed",
			})
		}
		return
	}

	// Handle /categories/{id}
	idStr := strings.TrimPrefix(path, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid category ID",
		})
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetCategory(w, r, id)
	case http.MethodPut:
		h.UpdateCategory(w, r, id)
	case http.MethodDelete:
		h.DeleteCategory(w, r, id)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Message: "Method not allowed",
		})
	}
}
