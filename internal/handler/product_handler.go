package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func ProductShow(c *gin.Context) {
	userID, existsID := c.Get("userID")
	userRole, existsRole := c.Get("userRole")
	if !existsID || !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	tmpl, err := template.ParseFiles("internal/templates/product.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки шаблона"})
		return
	}

	uid, _ := userID.(int)
	role, _ := userRole.(int)

	data := map[string]any{
		"UserID":   uid,
		"IsAdmin":  role == models.RoleAdmin,
		"IsSeller": role == models.RoleSeller,
		"CanBuy":   role == models.RoleClient,
	}

	if err := tmpl.Execute(c.Writer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "template error"})
		return
	}
}

func Product(c *gin.Context) {

	userID, existsID := c.Get("userID")
	userRole, existsRole := c.Get("userRole")
	if !existsID || !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	idStr := c.Query("id")
	idd, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	uid, _ := userID.(int)
	role, _ := userRole.(int)

	product, err := service.GetProdPoID(idd, role, uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	product.Category, err = repository.GetCategoryID(product.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if userRole == models.RoleClient {
		go func() {
			for _, offer := range product.Offers {
				_ = repository.TrackEvent(models.AnalyticsEvent{
					UserID:    &uid,
					ProductID: &product.ID,
					SellerID:  &offer.SellerID,
					EventType: models.EventView,
					Quantity:  1,
				})
			}
		}()
	}
	c.JSON(http.StatusOK, product)
}
