package handler

import (
	"html/template"
	"net/http"
	"strings"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"
	"tipodikayayagoda/internal/service"
	"tipodikayayagoda/internal/utils"

	"github.com/gin-gonic/gin"
)

func LoginShow(c *gin.Context) {

	tmpl, err := template.ParseFiles("internal/templates/login.html")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	tmpl.Execute(c.Writer, nil)
}

func Login(c *gin.Context) {

	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Login = strings.TrimSpace(req.Login)
	req.Password = strings.TrimSpace(req.Password)
	token, err := service.Login(req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("tokenn", token, 3600, "/", "", false, true)

	cookie, err := c.Cookie("tokenn")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	claims, err := utils.ValidateToken(cookie)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Невалидный токен"})
		c.Abort()
		return
	}
	idFloat, ok := claims["id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неверный формат ID в токене"})
		c.Abort()
		return
	}
	idint := int(idFloat)
	go repository.TrackEvent(models.AnalyticsEvent{
		UserID:    &idint,
		ProductID: nil,
		SellerID:  nil,
		EventType: models.EventLogin,
		Quantity:  1,
	})
	c.Redirect(http.StatusFound, "/index")
}
