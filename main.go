package main

import (
	controllers "github.com/Julia-Marcal/reusable-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("users/:id", controllers.GetUser)
	router.GET("users/", controllers.GetAllUsers)
	router.POST("users/", controllers.CreateUser)
	router.Run()
}
