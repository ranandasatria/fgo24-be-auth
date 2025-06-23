package controllers

import (
	"backend/models"
	"backend/responses"
	"fmt"
	"math/rand"
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

	models.Users = append(models.Users, &newUser)

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

func (a *AuthController) ForgotPassword(ctx *gin.Context) {
	type InputEmail struct {
		Email string `json:"email" form:"email"`
	}

	var input InputEmail

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	found := false
	for _, user := range models.Users {
		if user.Email == input.Email {
			found = true
			break
		}
	}
	if !found {
		ctx.JSON(http.StatusNotFound, responses.Response{
			Success: false,
			Message: "Email not registered",
		})
		return
	}

	otp := rand.Intn(1000000)

	models.ListForgotPass = append(models.ListForgotPass, &models.ForgotPass{
		Email:   input.Email,
		OTPCode: otp,
	})

	fmt.Printf("OTP for %s is: %06d\n", input.Email, otp)

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "OTP has been sent (check terminal)",
	})
}

func (a *AuthController) ResetPassword(ctx *gin.Context) {
	var input models.Reset

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	var foundOTP *models.ForgotPass
	for _, item := range models.ListForgotPass {
		if item.Email == input.Email && item.OTPCode == input.OTP {
			foundOTP = item
			break
		}
	}

	if foundOTP == nil {
		ctx.JSON(http.StatusUnauthorized, responses.Response{
			Success: false,
			Message: "OTP or email invalid",
		})
		return
	}

	if input.NewPass != input.ConfPass {
		ctx.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "Passwords do not match",
		})
		return
	}

	for _, user := range models.Users {
		if user.Email == input.Email {
			user.Password = input.NewPass
			break
		}
	}

	for i, item := range models.ListForgotPass {
		if item.Email == input.Email {
			models.ListForgotPass = append(models.ListForgotPass[:i], models.ListForgotPass[i+1:]...)
			break
		}
	}

	ctx.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "Password updated successfully",
	})
}
