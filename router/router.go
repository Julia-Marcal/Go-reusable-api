package router

import (
	controllers "github.com/Julia-Marcal/reusable-api/controllers"
	middlewares "github.com/Julia-Marcal/reusable-api/middlewares"
	metrics "github.com/Julia-Marcal/reusable-api/services/metrics"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	rateLimiter := middlewares.RateLimiting()

	api := router.Group("/api")
	{
		api.GET("login/", rateLimiter, controllers.GenerateToken)
		api.GET("/metrics", rateLimiter, metrics.PrometheusHandler())
		api.POST("users/", rateLimiter, controllers.CreateUser)
		authorized := api.Group("/v1/").Use(middlewares.Auth())
		{
			authorized.GET("user/", rateLimiter, controllers.GetUser)
			authorized.GET("users/", rateLimiter, controllers.GetAllUsers)
			authorized.DELETE("users/:id", rateLimiter, controllers.DeleteUser)
		}
	}
	router.Run()
}
