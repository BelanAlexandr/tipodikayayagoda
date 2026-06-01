package handler

import (
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func UploadImageHandler(c *gin.Context) {

	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	header, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image is required"})
		return
	}

	file, err := header.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open image"})
		return
	}
	defer file.Close()

	url, err := service.UploadImage(idd, file, header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}
