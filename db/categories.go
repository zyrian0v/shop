package db

func AddCategory(c Category) error {
	stmt := `INSERT INTO categories (name, slug)
	VALUES (?, ?)`
	_, err := handle.Exec(stmt, c.Name, c.Slug)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() (cs []Category, err error) {
	stmt := `SELECT id, name, slug FROM categories`
	rows, err := handle.Query(stmt)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		c := Category{}
		err = rows.Scan(&c.Id, &c.Name, &c.Slug)
		if err != nil {
			return
		}
		cs = append(cs, c)
	}
	return
}

func GetCategoryBySlug(slug string) (c Category, err error) {
	stmt := `SELECT name, slug FROM categories WHERE slug = ?`
	err = handle.QueryRow(stmt, slug).Scan(&c.Name, &c.Slug)
	return
}

func EditCategory(slug string, p Category) error {
	stmt := `UPDATE categories
	SET name = ?,
	slug = ?
	WHERE slug = ?`
	_, err := handle.Exec(stmt, p.Name, p.Slug, slug)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(slug string) error {
	stmt := `DELETE FROM categories
	WHERE slug = ?`
	_, err := handle.Exec(stmt, slug)
	if err != nil {
		return err
	}
	return nil
}
