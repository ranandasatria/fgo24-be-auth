package routers

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	authController := controllers.AuthController{}

	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)
	r.POST("/forgot-password", authController.ForgotPassword)
	r.POST("/reset-password", authController.ResetPassword)
}
