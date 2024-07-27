package controllers

import (
	"GO-JWT-Auth/initializers"
	"GO-JWT-Auth/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func FetchUserDetails(c *gin.Context, logger *zap.SugaredLogger) {

	id, _ := c.Get("id")

	logger.Info("User ID: ", id)

	var user models.Auth
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		logger.Info("User not found")
		return
	}

	logger.Info("User fetched successfully")

	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})

	logger.Info("id: ", user.ID, " email: ", user.Email)

}
