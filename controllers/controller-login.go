package controllers

import (
	"net/http"

	queries "github.com/Julia-Marcal/reusable-api/database/queries"
	security "github.com/Julia-Marcal/reusable-api/services/security"
	"github.com/gin-gonic/gin"
)

type PasswordInput struct {
	Password string `json:"password"`
}

func GetPass(c *gin.Context) {
	userId := c.Param("id")
	var passwordInput PasswordInput

	if err_pass := c.ShouldBindJSON(&passwordInput); err_pass != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	pass, err := queries.GetPassword(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}


	passwordCorrect, err := security.LoginSystem(passwordInput.Password, pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := gin.H{
		"message": "Password is correct",
		"login":   passwordCorrect,
	}

	c.JSON(http.StatusOK, response)
}
