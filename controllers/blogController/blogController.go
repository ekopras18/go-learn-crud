package blogController

import (
	"go-learn-crud-mysql/entities"
	"go-learn-crud-mysql/models/blogModel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	blogs := blogModel.Get()

	data := map[string]any{
		"blogs": blogs,
	}

	tmpl, err := template.ParseFiles("views/blog/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/blog/create.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	}

	if r.Method == "POST" {
		var blog entities.Blog

		blog.Title = r.FormValue("title")
		blog.Date = time.Now()
		blog.Author = r.FormValue("author")
		blog.Tags = r.FormValue("tags")
		blog.Content = []byte(r.FormValue("content"))
		blog.CreatedAt = time.Now()
		blog.UpdatedAt = time.Now()

		if ok := blogModel.Store(blog); !ok {
			tmpl, _ := template.ParseFiles("views/blog/create.html")
			tmpl.Execute(w, nil)
		}

		http.Redirect(w, r, "/blog", http.StatusSeeOther)

	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/blog/edit.html")
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

		blog := blogModel.Show(id)

		data := map[string]any{
			"blog": blog,
		}

		tmpl.Execute(w, data)
	}

	if r.Method == "POST" {
		var blog entities.Blog

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		// category.Id = id
		blog.Title = r.FormValue("title")
		blog.Date = time.Now()
		blog.Author = r.FormValue("author")
		blog.Tags = r.FormValue("tags")
		blog.Content = []byte(r.FormValue("content"))
		blog.UpdatedAt = time.Now()

		if ok := blogModel.Update(id, blog); !ok {
			// tmpl, _ := template.ParseFiles("views/category/edit.html")
			// tmpl.Execute(w, nil)
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/blog", http.StatusSeeOther)

	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	if err := blogModel.Destroy(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/blog", http.StatusSeeOther)
}
