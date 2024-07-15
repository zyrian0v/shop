package db

type Image struct {
	Id int
	Filename  string
	ProductId int
}

func AddImage(i Image) error {
	stmt := `INSERT INTO images (filename, product_id)
	VALUES (?, ?)`
	_, err := handle.Exec(stmt, i.Filename, i.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func GetProductImages(id int) (is []Image, err error) {
	stmt := `SELECT id, filename, product_id FROM images
	WHERE product_id = ?`
	rows, err := handle.Query(stmt, id)
	if err != nil {
		return
	}
	
	for rows.Next() {
		i := Image{}
		err = rows.Scan(&i.Id, &i.Filename, &i.ProductId)
		if err != nil {
			return
		}
		is = append(is, i)
	}
	return
}

func GetImageById(id int) (i Image, err error) {
	stmt := `SELECT filename, product_id FROM images
	WHERE id = ?`
	err = handle.QueryRow(stmt, id).Scan(&i.Filename, &i.ProductId)
	return
}

func DeleteImage(id int) error {
	stmt := `DELETE FROM images
	WHERE id = ?`
	_, err := handle.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}