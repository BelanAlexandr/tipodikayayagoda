package utils

import (
	"os"
	"strings"
)

func DeleteFileByURL(url string) {
	if url == "" {
		return
	}

	path := strings.TrimPrefix(url, "/")
	_ = os.Remove(path)
}
