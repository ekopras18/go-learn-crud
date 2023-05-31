package homeController

import (
	"go-learn-crud-mysql/utility/baseUtility"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Data yang ingin Anda tampilkan di konten
	data := struct {
		Title string
		Text  string
	}{
		Title: "Selamat datang",
		Text:  "Halaman Utama CRUD x Mysql.",
	}

	tmpl, err := template.ParseFiles("views/layout/base.html", "views/index.html")
	baseUtility.StatusInternalServer(w, err)

	tmpl.Execute(w, data)
}
