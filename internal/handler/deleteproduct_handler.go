package handler

import (
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/service"
	"tipodikayayagoda/internal/storage"

	"github.com/gin-gonic/gin"
)

func DeleteProductHandler(c *gin.Context) {
	userRoleValue, existsRole := c.Get("userRole")
	if !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	userrole, ok := userRoleValue.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		с.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = service.DeleteProd(storage.MinioClient, id, userrole)
	if err != nil {
		с.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messgae": "deleted"})
}
