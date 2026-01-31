package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/database"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/handlers"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/repositories"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/services"
	"github.com/spf13/viper"
)

// Config struct buat nyimpen konfigurasi aplikasi
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	// ==================== LOAD CONFIG ====================
	// Load environment variables pake Viper
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Cek kalo ada file .env, baca dari situ
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	// Ambil config dari environment
	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Default port kalo gak diset
	if config.Port == "" {
		config.Port = "8080"
	}

	// ==================== SETUP DATABASE ====================
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Gagal connect ke database:", err)
	}
	defer db.Close()

	// ==================== DEPENDENCY INJECTION ====================
	// Ini kayak Manager di restoran yang kenalin Pelayan, Koki, sama Anak Gudang

	// Product: Repository -> Service -> Handler
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Category: Repository -> Service -> Handler
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// ==================== SETUP ROUTES ====================
	// Product routes
	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)

	// Category routes
	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// ==================== START SERVER ====================
	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server jalan di", addr)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET/POST   /api/products")
	fmt.Println("  GET/PUT/DELETE /api/products/{id}")
	fmt.Println("  GET/POST   /api/categories")
	fmt.Println("  GET/PUT/DELETE /api/categories/{id}")
	fmt.Println("  GET        /health")

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Waduh, server gagal jalan:", err)
	}
}
