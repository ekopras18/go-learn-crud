package categoriesController

import (
	"go-learn-crud/entities"
	"go-learn-crud/models/categoryModel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categoryModel.Get()

	data := map[string]any{
		"categories": categories,
	}

	tmpl, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categoryModel.Store(category); !ok {
			tmpl, _ := template.ParseFiles("views/category/create.html")
			tmpl.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)

	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			// return
			panic(err)
		}

		category := categoryModel.Show(id)

		data := map[string]any{
			"category": category,
		}

		tmpl.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.Category

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		// category.Id = id
		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if ok := categoryModel.Update(id, category); !ok {
			// tmpl, _ := template.ParseFiles("views/category/edit.html")
			// tmpl.Execute(w, nil)
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)

	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	if err := categoryModel.Destroy(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
