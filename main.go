package main

import (
	"go-learn-crud/config"
	"go-learn-crud/controllers/categoriesController"
	"go-learn-crud/controllers/homeController"
	"log"
	"net/http"
)

func main() {

	// Initialize database connection
	config.ConnectDB()

	// Initialize routes
	http.HandleFunc("/", homeController.Index)

	//

	http.HandleFunc("/categories", categoriesController.Index)
	http.HandleFunc("/categories/create", categoriesController.Create)
	http.HandleFunc("/categories/edit", categoriesController.Edit)
	http.HandleFunc("/categories/delete", categoriesController.Delete)

	// Run the server
	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
