package routers

import "github.com/gin-gonic/gin"

func CombineRouter(r *gin.Engine){
	authRouter(r.Group("/auth"))
	authRouter(r.Group("/users"))
}