package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/database"
	_ "github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/docs"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/handlers"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/repositories"
	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/services"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Kasir API
// @version 1.0
// @description API untuk sistem kasir dengan manajemen produk dan kategori
// @host localhost:8080
// @BasePath /

// Config
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	// Load config dengan Viper
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Dependency Injection
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup routes
	http.HandleFunc("/api/products", productHandler.HandleProducts)
	http.HandleFunc("/api/products/", productHandler.HandleProductByID)

	http.HandleFunc("/api/categories", categoryHandler.HandleCategories)
	http.HandleFunc("/api/categories/", categoryHandler.HandleCategoryByID)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// Swagger docs
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Start server
	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)
	fmt.Println("Swagger: http://localhost:" + config.Port + "/swagger/index.html")

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}
