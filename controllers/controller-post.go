package controllers

import (
	"net/http"

	database "github.com/Julia-Marcal/reusable-api/database"
	queries "github.com/Julia-Marcal/reusable-api/database/queries"
	cache "github.com/Julia-Marcal/reusable-api/repository/cache/caching-func"
	validation "github.com/Julia-Marcal/reusable-api/services/validation"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user database.User

	validated := validation.UserValidator(user)

	if err := c.ShouldBindJSON(&user); err != nil || validated {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	CacheErr := cache.CacheUser(user)
	if CacheErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cache user",
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
