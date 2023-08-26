package controllers

import (
	"net/http"

	queries "github.com/Julia-Marcal/reusable-api/database/queries"
	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Email string `json:"email"`
}

func GetUser(c *gin.Context) {
	var request UserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	user, err := queries.FindUser(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User returned successfully",
		"user":    user,
	})
}
