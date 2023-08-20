package controllers

import (
	"net/http"

	queries "github.com/Julia-Marcal/reusable-api/database/queries"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, exists := c.Params.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	user, err := queries.FindOne(userId)
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

	c.JSON(http.StatusOK, user)

}
