package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hspgit/go_receipts/internal/database"
	"github.com/hspgit/go_receipts/internal/models"
)

// GetAllUsers returns all users from the database
func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := database.GetDB().Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"count": len(users),
	})
}
