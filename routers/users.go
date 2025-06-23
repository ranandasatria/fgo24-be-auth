package routers

import "github.com/gin-gonic/gin"

func userRouter(r *gin.RouterGroup){
	r.GET("/")
	r.GET("/:id")
	r.POST("/")
	r.PATCH("/:id")
	r.DELETE("/:id")
}