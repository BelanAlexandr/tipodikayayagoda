package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/utils"
)

func UploadImage(id int, file multipart.File, header *multipart.FileHeader) (string, error) {

	oldURL := repository.GetImageURL(id)
	if oldURL != "" {
		utils.DeleteFileByURL(oldURL)
	}
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), header.Filename)
	path := "./static/uploads/" + filename

	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}
	repository.Updateimage("/static/uploads/"+filename, id)
	return "/static/uploads/" + filename, nil
}
