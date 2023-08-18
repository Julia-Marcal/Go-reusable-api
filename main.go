package main

import (
	repository "github.com/Julia-Marcal/reusable-api/repository"
	_ "github.com/gin-gonic/gin"
)

func main() {
	repository.NewPostgres()
	/*
		router := gin.Default()
		router.GET("/ping", controllers.Pong)
		router.Run()
	*/
}
