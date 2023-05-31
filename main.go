package main

import (
	"fmt"
	"go-learn-crud-mysql/config"
	"go-learn-crud-mysql/controllers/blogController"
	"go-learn-crud-mysql/controllers/categoriesController"
	"go-learn-crud-mysql/controllers/homeController"
	"go-learn-crud-mysql/utility/baseUtility"
	"log"
	"net/http"
)

func main() {

	// Initialize database connection
	config.ConnectDB()

	// Initialize routes
	http.HandleFunc("/", homeController.Index)

	//categories
	http.HandleFunc("/categories", categoriesController.Index)
	http.HandleFunc("/categories/create", categoriesController.Create)
	http.HandleFunc("/categories/edit", categoriesController.Edit)
	http.HandleFunc("/categories/delete", categoriesController.Delete)

	//blog
	http.HandleFunc("/blog", blogController.Index)
	http.HandleFunc("/blog/create", blogController.Create)
	http.HandleFunc("/blog/edit", blogController.Edit)
	http.HandleFunc("/blog/delete", blogController.Delete)

	// Run the server
	log.Println("Server started on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	baseUtility.Catch(err)
	fmt.Println("Server stopped.")

}
