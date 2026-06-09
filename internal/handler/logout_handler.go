package handler

import (
	"net/http"
	"tipodikayayagoda/internal/models"
	"tipodikayayagoda/internal/repository"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	userIdValue, exists := c.Get("userID")

	if exists {
		if id, ok := userIdValue.(int); ok {

			go repository.TrackEvent(models.AnalyticsEvent{
				UserID:    &id,
				ProductID: nil,
				SellerID:  nil,
				EventType: models.EventLogout,
				Quantity:  1,
			})

		}
	}

	c.SetCookie("tokenn", " ", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}
