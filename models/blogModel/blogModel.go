package blogModel

import (
	"go-learn-crud-mysql/config"
	"go-learn-crud-mysql/entities"
)

func Get() []entities.Blog {
	rows, err := config.DB.Query("SELECT * FROM blog")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var blogs []entities.Blog

	for rows.Next() {
		var blog entities.Blog

		if err := rows.Scan(&blog.Id, &blog.Title, &blog.Date, &blog.Author, &blog.Tags, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			panic(err)
		}

		blogs = append(blogs, blog)
	}

	return blogs

}

func Store(blog entities.Blog) bool {
	result, err := config.DB.Exec("INSERT INTO blog (title, date, tags, author, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)", blog.Title, blog.Date, blog.Tags, blog.Author, blog.Content, blog.CreatedAt, blog.UpdatedAt)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0

}

func Show(id int) entities.Blog {

	row := config.DB.QueryRow("SELECT * FROM blog WHERE id = ?", id)

	var blog entities.Blog
	if err := row.Scan(&blog.Id, &blog.Title, &blog.Date, &blog.Author, &blog.Tags, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
		panic(err)
	}

	return blog

}

func Update(id int, blog entities.Blog) bool {
	result, err := config.DB.Exec("UPDATE blog SET title = ?, author = ?, tags = ?, content = ?, updated_at = ? WHERE id = ?", blog.Title, blog.Author, blog.Tags, blog.Content, blog.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowsAffected > 0

}

func Destroy(id int) error {
	_, err := config.DB.Exec("DELETE FROM blog WHERE id = ?", id)

	return err
}
