package controllers

import (
	"net/http"
	queries "github.com/Julia-Marcal/reusable-api/database/queries"
	database "github.com/Julia-Marcal/reusable-api/database"
	"github.com/gin-gonic/gin"
)


func CreateUser(c *gin.Context) {
	var user database.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	err := queries.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}