package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/achmadhadikurnia/bootcamp-jago-golang-dasar/handlers"
)

func main() {
	// Initialize handlers
	categoryHandler := handlers.NewCategoryHandler()

	// Register routes
	http.Handle("/categories", categoryHandler)
	http.Handle("/categories/", categoryHandler)

	// Start server
	port := ":8080"
	fmt.Printf("🚀 Server is running on http://localhost%s\n", port)
	fmt.Println("📚 Available endpoints:")
	fmt.Println("   GET    /categories      - Get all categories")
	fmt.Println("   POST   /categories      - Create a category")
	fmt.Println("   GET    /categories/{id} - Get a category by ID")
	fmt.Println("   PUT    /categories/{id} - Update a category by ID")
	fmt.Println("   DELETE /categories/{id} - Delete a category by ID")
	fmt.Println()

	log.Fatal(http.ListenAndServe(port, nil))
}
