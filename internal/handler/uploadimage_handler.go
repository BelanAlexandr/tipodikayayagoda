package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"tipodikayayagoda/internal/service"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/api/uploadimage/")
	idd, err := strconv.Atoi(id)
	fmt.Println(id)
	if err != nil {
		http.Error(w, "invalid product ID", 400)
		return
	}
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "image is required", 400)
		return
	}
	defer file.Close()
	fmt.Println("aaa")
	url, err := service.UploadImage(idd, file, header)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"url": url,
	})
}
