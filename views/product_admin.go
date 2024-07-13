package views

import (
	"fmt"
	slugify "github.com/gosimple/slug"
	"html/template"
	"net/http"
	"shop/db"
	"strconv"
)

type Admin struct{}

func (v Admin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/base.html",
		"templates/admin.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}

type EditProductList struct {
	Products []db.Product
}

func (v EditProductList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ps, err := db.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	v.Products = ps

	files := []string{
		"templates/base.html",
		"templates/edit_product_list.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}

type EditProduct struct {
	db.Product
	Categories []db.Category
}

func (v EditProduct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		v.post(w, r)
		return
	}

	slug := r.PathValue("slug")
	p, err := db.GetProductBySlug(slug)
	if err != nil {
		fmt.Println(err)
		http.NotFound(w, r)
		return
	}
	v.Product = p

	cs, err := db.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	v.Categories = cs

	files := []string{
		"templates/base.html",
		"templates/edit_product.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}

func (v EditProduct) post(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	catid, err := strconv.Atoi(r.FormValue("category"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	p := db.Product{
		Name:   r.FormValue("name"),
		Slug:   r.FormValue("slug"),
		Detail: r.FormValue("detail"),
		CategoryId: catid,
	}
	err = db.EditProduct(slug, p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/admin/products/edit", 303)
}

type NewProduct struct{
	Categories []db.Category
}

func (v NewProduct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		v.post(w, r)
		return
	}

	cs, err := db.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	v.Categories = cs

	files := []string{
		"templates/base.html",
		"templates/new_product.html",
	}
	tmpl := template.Must(template.ParseFiles(files...))
	if err := tmpl.Execute(w, v); err != nil {
		fmt.Println(err)
	}
}

func (v NewProduct) post(w http.ResponseWriter, r *http.Request) {
	catid, err := strconv.Atoi(r.FormValue("category"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	p := db.Product{
		Name:   r.FormValue("name"),
		Slug:   slugify.Make(r.FormValue("name")),
		Detail: r.FormValue("detail"),
		CategoryId: catid,
	}

	_, err = db.AddProduct(p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/products/"+p.Slug, 303)
}

type DeleteProduct struct{}

func (v DeleteProduct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	err := db.DeleteProduct(slug)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, "/admin/products/edit", 303)
}
