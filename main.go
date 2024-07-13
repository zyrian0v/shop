package main

import (
	"flag"
	"fmt"
	"net/http"
	"shop/db"
	"shop/views"
)

func main() {
	schemaFlag := flag.Bool("schema", false, "apply schema")
	flag.Parse()

	handle := db.InitializeHandle()
	defer handle.Close()
	if *schemaFlag {
		db.ApplySchema()
		return
	}

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.Handle("/", views.Index{})
	http.Handle("/products/{slug}", views.ShowProduct{})
	http.Handle("/category/{slug}", views.ShowCategory{})
	http.Handle("/admin", views.Admin{})
	http.Handle("/admin/products/new", views.NewProduct{})
	http.Handle("/admin/products/edit", views.EditProductList{})
	http.Handle("/admin/products/edit/{slug}", views.EditProduct{})
	http.Handle("/admin/products/delete/{slug}", views.DeleteProduct{})

	http.Handle("/admin/categories/new", views.NewCategory{})
	http.Handle("/admin/categories/edit", views.EditCategoryList{})
	http.Handle("/admin/categories/edit/{slug}", views.EditCategory{})
	http.Handle("/admin/categories/delete/{slug}", views.DeleteCategory{})

	fmt.Println("serving on :8080")
	http.ListenAndServe(":8080", nil)
}
