package controllers

import (
	"net/http"

	database "github.com/Julia-Marcal/reusable-api/database"
	queries "github.com/Julia-Marcal/reusable-api/database/queries"
	auth "github.com/Julia-Marcal/reusable-api/services/auth"
	security "github.com/Julia-Marcal/reusable-api/services/security"
	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user database.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	pass, err := queries.CheckPassword(request.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err_pass := security.LoginSystem(request.Password, pass)
	if err_pass != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokenString, err := auth.GenerateJWT(user.Email, user.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
