package db

type Image struct {
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
	stmt := `SELECT filename, product_id FROM images
	WHERE product_id = ?`
	rows, err := handle.Query(stmt, id)
	if err != nil {
		return
	}
	
	for rows.Next() {
		i := Image{}
		err = rows.Scan(&i.Filename, &i.ProductId)
		if err != nil {
			return
		}
		is = append(is, i)
	}
	return
}
