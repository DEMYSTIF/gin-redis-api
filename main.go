package main

import (
	"github.com/DEMYSTIF/gin-redis-api/config"
	"github.com/DEMYSTIF/gin-redis-api/handlers"
	"github.com/DEMYSTIF/gin-redis-api/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	client := config.ConnectRedis()

	router := gin.Default()
	router.Use(middlewares.Authority())
	router.POST("/create", func(ctx *gin.Context) {
		handlers.CreateOne(ctx, client)
	})
	router.GET("/read/:id", func(ctx *gin.Context) {
		handlers.ReadOne(ctx, client)
	})
	router.Run("localhost:8080")
}