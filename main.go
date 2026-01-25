package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/handlers"
)

func main() {
	// Initialize handlers
	categoryHandler := handlers.NewCategoryHandler()

	// Register routes
	http.Handle("/categories", categoryHandler)
	http.Handle("/categories/", categoryHandler)

	// Get port from environment variable (Railway compatible)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port for local development
	}

	// Start server
	fmt.Printf("🚀 Server is running on port %s\n", port)
	fmt.Println("📚 Available endpoints:")
	fmt.Println("   GET    /categories      - Get all categories")
	fmt.Println("   POST   /categories      - Create a category")
	fmt.Println("   GET    /categories/{id} - Get a category by ID")
	fmt.Println("   PUT    /categories/{id} - Update a category by ID")
	fmt.Println("   DELETE /categories/{id} - Delete a category by ID")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
