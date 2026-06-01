package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/service"

	"github.com/gin-gonic/gin"
)

func IndexHandlerShow(c *gin.Context) {
	userIDValue, existsID := c.Get("userID")
	userRoleValue, existsRole := c.Get("userRole")
	if !existsID || !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}
	userrole, ok := userRoleValue.(int)
	userID, ok := userIDValue.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID пользователя"})
		return
	}
	tmpl, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	tmpl.Execute(c.Writer, map[string]any{
		"UserID":        userID,
		"IsAdmin":       userrole == models.Roles.AdminID,
		"IsSeller":      userrole == models.Roles.SellerID,
		"CanBuy":        userrole == models.Roles.ClientID,
		"CanAddUser":    userrole == models.Roles.AdminID,
		"CanAddProduct": userrole == models.Roles.AdminID || userrole == models.Roles.SellerID,
	})
}

type ProductsResponse struct {
	Products   []models.Product `json:"products"`
	TotalCount int              `json:"totalCount"`
}

func IndexHandler(c *gin.Context) {
	userIDValue, existsID := c.Get("userID")
	userRoleValue, existsRole := c.Get("userRole")
	if !existsID || !existsRole {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Данные авторизации не найдены"})
		return
	}

	userID, ok1 := userIDValue.(int)
	userRole, ok2 := userRoleValue.(int)
	if !ok1 || !ok2 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат ID или роли пользователя"})
		return
	}

	search := c.Query("search")
	sort := c.Query("sort")

	categoryID, err := strconv.Atoi(c.DefaultQuery("category", "0"))
	if err != nil || categoryID < 0 {
		categoryID = 0
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "6"))
	if err != nil || limit < 1 {
		limit = 6
	}

	products, totalCount, err := service.GetProducts(userRole, userID, search, page, limit, sort, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error loading products"})
		return
	}

	response := ProductsResponse{
		Products:   products,
		TotalCount: totalCount,
	}
	c.JSON(http.StatusOK, response)
}
