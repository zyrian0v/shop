package views

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func saveImage(fh *multipart.FileHeader) (string, error) {
	f, err := fh.Open()
	if err != nil {
		return "", err
	}
	ut := strconv.Itoa(int(time.Now().Unix()))
	fname := fmt.Sprintf("%v_%v", ut, fh.Filename)
	out, err := os.Create("images/"+fname)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, f)
	if err != nil {
		return "", err
	}
	return fname, nil
}
