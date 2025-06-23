package main

import (
	"backend/controllers"
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authController := controllers.AuthController{}
	var authHandler handlers.AuthHandler = &authController

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Backend is running"})
	})

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	r.Run()
}

// routers.CombineRouter