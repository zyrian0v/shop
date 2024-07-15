package views

import (
	"fmt"
	"html/template"
	"net/http"
	"shop/db"
)

type Index struct {
	Products   []db.ProductWithImage
	Categories []db.Category
}

func (v Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ps, err := db.GetAllProductsWithImage()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	v.Products = ps

	cs, err := db.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	v.Categories = cs

	files := []string{
		"templates/base.html",
		"templates/index.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}

type ShowProduct struct {
	db.Product
	Images []db.Image
}

func (v ShowProduct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	p, err := db.GetProductBySlug(slug)
	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	v.Product = p

	is, err := db.GetProductImages(p.Id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	v.Images = is

	files := []string{
		"templates/base.html",
		"templates/show_product.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}

type ShowCategory struct {
	Products []db.ProductWithImage
}

func (v ShowCategory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	ps, err := db.GetProductsByCategory(slug)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	v.Products = ps

	files := []string{
		"templates/base.html",
		"templates/show_category.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}
