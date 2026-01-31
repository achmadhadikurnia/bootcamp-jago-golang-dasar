package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Product itu struct buat nyimpen data produk di sistem kasir kita
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

// Ini storage sementara di RAM, nanti bakal diganti pake database beneran
var products = []Product{
	{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
	{ID: 3, Name: "kecap", Price: 12000, Stock: 20},
}

// getProductByID buat handle GET /api/products/{id}
func getProductByID(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari URL path
	// Contoh: /api/products/123 -> dapet ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Cari product yang ID-nya cocok
	for _, p := range products {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Kalo sampe sini berarti gak ketemu
	http.Error(w, "Product not found", http.StatusNotFound)
}

// updateProduct buat handle PUT /api/products/{id}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")

	// Convert jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ambil data dari body request
	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Loop semua products, cari ID yg cocok, terus update datanya
	for i := range products {
		if products[i].ID == id {
			updatedProduct.ID = id
			products[i] = updatedProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// deleteProduct buat handle DELETE /api/products/{id}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dulu
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")

	// Convert jadi int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Loop products, cari ID yg mau dihapus
	for i, p := range products {
		if p.ID == id {
			// Bikin slice baru tanpa data yg mau dihapus
			products = append(products[:i], products[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Product deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	// Route buat akses product pake ID
	// GET localhost:8080/api/products/{id}
	// PUT localhost:8080/api/products/{id}
	// DELETE localhost:8080/api/products/{id}
	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProductByID(w, r)
		} else if r.Method == "PUT" {
			updateProduct(w, r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}
	})

	// Route buat ambil semua products atau bikin product baru
	// GET localhost:8080/api/products
	// POST localhost:8080/api/products
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products)
		} else if r.Method == "POST" {
			// Baca data dari request body
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// Masukin data ke variable products
			newProduct.ID = len(products) + 1
			products = append(products, newProduct)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(newProduct)
		}
	})

	// Health check endpoint buat ngecek server masih idup apa enggak
	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	fmt.Println("Server jalan di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Waduh, server gagal jalan nih")
	}
}
