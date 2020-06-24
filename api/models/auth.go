package models

type RegisterUser struct {
	Username  string `json:"username" form:"username"  binding:"required,min=6,max=50"`
	Password  string `json:"password" form:"password"  binding:"required,min=6,max=20"`
	FirstName string `json:"firstName" form:"firstName"  binding:"required"`
	LastName  string `json:"lastName" form:"lastName"  binding:"required"`
	Email     string `json:"email" form:"email"  binding:"required,email"`
}
