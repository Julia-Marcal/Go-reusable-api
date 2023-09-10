package router

import (
	controllers "github.com/Julia-Marcal/reusable-api/controllers"
	middlewares "github.com/Julia-Marcal/reusable-api/middlewares"
	metrics "github.com/Julia-Marcal/reusable-api/services/metrics"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("login/", controllers.GenerateToken)
		api.GET("/metrics", metrics.PrometheusHandler())
		authorized := api.Group("/v1/").Use(middlewares.Auth())
		{
			authorized.GET("user/", controllers.GetUser)
			authorized.GET("users/", controllers.GetAllUsers)
			authorized.DELETE("users/:id", controllers.DeleteUser)
			authorized.POST("users/", controllers.CreateUser)
		}
	}
	router.Run()
}
