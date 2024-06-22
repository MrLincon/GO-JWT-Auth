package controllers

import (
	"GO-JWT-Auth/initializers"
	"GO-JWT-Auth/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FetchUserDetails(c *gin.Context) {

	id, _ := c.Get("id")

	var user models.Auth
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}
