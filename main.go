package main

import (
	"backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Backend is running"})
	})

	routers.CombineRouter(r)

	r.Run()
}
