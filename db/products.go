package db

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

func GetProductBySlug(slug string) (p Product, err error) {
	stmt := `SELECT name, slug, detail, category_id FROM products
	WHERE slug = ?`
	err = handle.QueryRow(stmt, slug).Scan(&p.Name, &p.Slug, &p.Detail, &p.CategoryId)
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

func GetProductsByCategory(slug string) (ps []Product, err error){
	stmt := `SELECT id FROM categories WHERE slug = ?`
	var catid int
	err = handle.QueryRow(stmt, slug).Scan(&catid)
	if err != nil {
		return
	}

	stmt = `SELECT name, slug, detail, category_id
	FROM products WHERE category_id = ?`
	rows, err := handle.Query(stmt, catid)
	if err != nil {
		return
	}
	for rows.Next() {
		p := Product{}
		err = rows.Scan(&p.Name, &p.Slug, &p.Detail, &p.CategoryId)
		if err != nil {
			return
		}
		ps = append(ps, p)
	}
	return

}
