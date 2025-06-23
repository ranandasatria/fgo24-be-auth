package controllers

import (
	"backend/models"
	"backend/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (a *AuthController) Register(ctx *gin.Context) {
	var newUser models.User

	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	for _, u := range models.Users {
		if u.Email == newUser.Email {
			ctx.JSON(http.StatusConflict, responses.Response{
				Success: false,
				Message: "Email already registered",
			})
			return
		}
	}

	models.Users = append(models.Users, newUser)

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Register success",
		Results: newUser,
	})
}

func (a *AuthController) Login(ctx *gin.Context) {
	var loginData models.User

	if err := ctx.ShouldBind(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	for _, u := range models.Users {
		if u.Email == loginData.Email && u.Password == loginData.Password {
			ctx.JSON(http.StatusOK, responses.Response{
				Success: true,
				Message: "Login success",
				Results: u,
			})
			return
		}
	}

	ctx.JSON(http.StatusUnauthorized, responses.Response{
		Success: false,
		Message: "Email or password incorrect",
	})
}
