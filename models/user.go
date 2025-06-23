package models

type User struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

var Users []*User

type ForgotPass struct {
	Email   string
	OTPCode int
}

var ListForgotPass []*ForgotPass

type Reset struct {
	Email    string `json:"email" form:"email"`
	OTP      int    `json:"otp" form:"otp"`
	NewPass  string `json:"newPass" form:"newPass"`
	ConfPass string `json:"confPass" form:"confPass"`
}
