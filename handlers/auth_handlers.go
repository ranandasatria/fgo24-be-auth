package handlers

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}
