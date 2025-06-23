package models

type User struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var Users []User
