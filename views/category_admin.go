package views

import (
	"fmt"
	"net/http"
	"html/template"
	"shop/db"
	slugify "github.com/gosimple/slug"
)

type NewCategory struct{}

func (v NewCategory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		v.post(w, r)
		return
	}

	files := []string{
		"templates/base.html",
		"templates/new_category.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}

func (v NewCategory) post(w http.ResponseWriter, r *http.Request) {
	c := db.Category{
		Name:   r.FormValue("name"),
		Slug:   slugify.Make(r.FormValue("name")),
	}
	err := db.AddCategory(c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/admin", 303)
}

type EditCategoryList struct {
	Categories []db.Category
}

func (v EditCategoryList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cs, err := db.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	v.Categories = cs

	files := []string{
		"templates/base.html",
		"templates/edit_category_list.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}

type EditCategory struct {
	db.Category
}

func (v EditCategory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		v.post(w, r)
		return
	}

	slug := r.PathValue("slug")
	c, err := db.GetCategoryBySlug(slug)
	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	v.Category = c

	files := []string{
		"templates/base.html",
		"templates/edit_category.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}

func (v EditCategory) post(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	c := db.Category{
		Name: r.FormValue("name"),
		Slug: r.FormValue("slug"),
	}
	err := db.EditCategory(slug, c)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/admin/categories/edit", 303)
}

type DeleteCategory struct{}

func (v DeleteCategory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	err := db.DeleteCategory(slug)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/admin/categories/edit", 303)
}