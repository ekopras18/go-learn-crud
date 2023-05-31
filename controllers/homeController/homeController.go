package homeController

import (
	"go-learn-crud-mysql/models/blogModel"
	"go-learn-crud-mysql/utility/baseUtility"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	blogs := blogModel.Get()

	// Data yang ingin Anda tampilkan di konten
	data := map[string]any{
		"blogs":       blogs,
		"title":       "Welcome to our blog!",
		"page_tittle": "This is a simple CRUD application using Go and MySQL.",
		"page_active": "blog",
	}

	tmpl, err := template.ParseFiles("views/layout/base.html", "views/index.html")
	baseUtility.StatusInternalServer(w, err)

	tmpl.Execute(w, data)
}
