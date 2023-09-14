package controllers

import (
	"net/http"

	queries "github.com/Julia-Marcal/reusable-api/database/queries"
	cache "github.com/Julia-Marcal/reusable-api/repository/cache/caching-func"
	validation "github.com/Julia-Marcal/reusable-api/services/validation"
	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUser(c *gin.Context) {
	var request UserRequest

	validated := validation.EmailPassValidator(request)

	if err := c.ShouldBindJSON(&request); err != nil || validated {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	cachedUser, cacheErr := cache.GetCachedUser(request.Email)

	if cachedUser.Id != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "User retrieved from cache",
			"user":    cachedUser,
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

	cacheErr = cache.CacheUser(*user)
	if cacheErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cache user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User returned successfully",
		"user":    user,
	})
}
