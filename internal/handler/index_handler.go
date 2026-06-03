package handler

import (
	"fmt"
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
	fmt.Print(userIDValue)
	userrole, ok := userRoleValue.(int)
	userID, ok := userIDValue.(int)
	fmt.Println(userrole, userID)
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
		"IsAdmin":       userrole == models.RoleAdmin,
		"IsSeller":      userrole == models.RoleSeller,
		"CanBuy":        userrole == models.RoleClient,
		"CanAddUser":    userrole == models.RoleAdmin,
		"CanAddProduct": userrole == models.RoleAdmin || userrole == models.RoleSeller,
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

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	lastID, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
	if err != nil || lastID < 0 {
		lastID = 0
	}

	lastPrice, err := strconv.ParseFloat(c.DefaultQuery("last_price", "0.0"), 64)
	if err != nil || lastPrice < 0 {
		lastPrice = 0.0
	}

	lastRank, err := strconv.ParseFloat(c.DefaultQuery("last_rank", "0.0"), 64)
	if err != nil || lastRank < 0 {
		lastRank = 0.0
	}

	lastLength, err := strconv.Atoi(c.DefaultQuery("last_length", "0"))
	if err != nil || lastLength < 0 {
		lastLength = 0
	}

	products, totalCount, err := service.GetProducts(userRole, userID, search, lastID, lastPrice, lastRank, lastLength, limit, sort, categoryID)
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
