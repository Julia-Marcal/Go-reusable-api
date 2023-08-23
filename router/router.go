package router

import (
	controllers "github.com/Julia-Marcal/reusable-api/controllers"
	"github.com/gin-gonic/gin"
)

func Start_Router() {
	router := gin.Default()
	router.GET("users/:id", controllers.GetUser)
	router.GET("login/:id", controllers.GetPass)
	router.DELETE("users/:id", controllers.DeleteUser)
	router.GET("users/", controllers.GetAllUsers)
	router.POST("users/", controllers.CreateUser)
	router.Run()
}
