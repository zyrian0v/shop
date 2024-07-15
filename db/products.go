package db

import (
	"database/sql"
)

type Product struct {
	Id int
	Name       string
	Slug       string
	Detail     string
	CategoryId int
}

type ProductWithImage struct {
	Product
	Filename string
}

func GetAllProducts() (ps []Product, err error) {
	stmt := `SELECT name, slug, detail FROM products`
	rows, err := handle.Query(stmt)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := Product{}
		err = rows.Scan(&p.Name, &p.Slug, &p.Detail)
		if err != nil {
			return
		}
		ps = append(ps, p)
	}
	return
}

func GetAllProductsWithImage() ([]ProductWithImage, error) {
	stmt := `SELECT id, name, slug, detail FROM products`
	rows, err := handle.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var ps []ProductWithImage
	for rows.Next() {		
		p := ProductWithImage{}
		err := rows.Scan(&p.Id, &p.Name, &p.Slug, &p.Detail)
		if err != nil {
			return nil, err
		}
		stmt2 := `SELECT filename FROM images WHERE product_id = ? LIMIT 1`
		err = handle.QueryRow(stmt2, p.Id).Scan(&p.Filename)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}

func GetProductsByCategory(catid int) ([]ProductWithImage, error) {
	stmt := `SELECT id, name, slug, detail
	FROM products WHERE category_id = ?`
	rows, err := handle.Query(stmt, catid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ps []ProductWithImage
	for rows.Next() {
		p := ProductWithImage{}
		err := rows.Scan(&p.Id, &p.Name, &p.Slug, &p.Detail)
		if err != nil {
			return nil, err
		}
		stmt2 := `SELECT filename FROM images WHERE product_id = ? LIMIT 1`
		err = handle.QueryRow(stmt2, p.Id).Scan(&p.Filename)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil

}

func GetProductBySlug(slug string) (p Product, err error) {
	stmt := `SELECT id, name, slug, detail, category_id FROM products
	WHERE slug = ?`
	err = handle.QueryRow(stmt, slug).Scan(&p.Id, &p.Name, &p.Slug, &p.Detail, &p.CategoryId)
	return
}

func GetProductId(slug string) (id int, err error) {
	stmt := `SELECT id FROM products WHERE slug = ?`
	err = handle.QueryRow(stmt, slug).Scan(&id)
	return
}

func AddProduct(p Product) (int64, error) {
	stmt := `INSERT INTO products (name, slug, detail, category_id)
	VALUES (?, ?, ?, ?)`
	res, err := handle.Exec(stmt, p.Name, p.Slug, p.Detail, p.CategoryId)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func EditProduct(slug string, p Product) error {
	stmt := `UPDATE products
	SET name = ?,
	slug = ?,
	detail = ?,
	category_id = ?
	WHERE slug = ?`
	_, err := handle.Exec(stmt, p.Name, p.Slug, p.Detail, p.CategoryId, slug)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(slug string) error {
	stmt := `DELETE FROM products
	WHERE slug = ?`
	_, err := handle.Exec(stmt, slug)
	if err != nil {
		return err
	}
	return nil
}
