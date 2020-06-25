package models

import (
	"strings"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" form:"username" gorm:"type:varchar(50);not null;unique_index" binding:"required,min=6,max=50"`
	Password  string `json:"password" form:"password" gorm:"not null" binding:"required,min=6,max=20"`
	FirstName string `json:"firstName" form:"firstName" gorm:"type:varchar(100);not null" binding:"required"`
	LastName  string `json:"lastName" form:"lastName" gorm:"type:varchar(100);not null" binding:"required"`
	Email     string `json:"email" form:"email" gorm:"type:varchar(100);unique_index" binding:"required,email"`
	Slug      string `json:"slug" form:"slug" uri:"slug"  gorm:"type:varchar(50);unique_index"`
	Status    string `json:"status" form:"status" sql:"type:enum('A','I');DEFAULT:'A'"`
	Avatar    string `json:"avatar" form:"avatar"`
	RoleID    int    `json:"roleId" form:"roleId"`

	Role Role `json:"role" `
}

type ChangePassword struct {
	CurrentPassword string `json:"current_password" form:"current_password" binding:"required" example:"password"` // รหัสผ่านปัจจุบัน
	NewPassword     string `json:"new_password" form:"new_password" binding:"required" example:"password123"`      // รหัสผ่านใหม่
}

type UploadAvatar struct {
	Avatar string `json:"avatar" form:"avatar" binding:"required" example:"avatar.png"` // รูปภาพ
}

func (u *User) BeforeSave(db *gorm.DB) (err error) {
	u.Slug = slug.Make(u.Username)
	return
}

func (u User) Serialize() map[string]interface{} {

	replaceAllFlag := -1

	return map[string]interface{}{
		"id":        u.ID,
		"username":  u.Username,
		"firstName": u.FirstName,
		"lastName":  u.LastName,
		"email":     u.Email,
		"slug":      u.Slug,
		"status":    u.Status,
		"avatar":    strings.Replace(u.Avatar, "\\", "/", replaceAllFlag),
		"roleId":    u.RoleID,
		"role":      u.Role.Serialize(),
	}
}

// GenerateJwtToken -- Generate JWT token associated to this user
func (u *User) GenerateJwtToken() string {
	// jwt.New(jwt.GetSigningMethod("HS512"))
	jwtToken := jwtgo.New(jwtgo.SigningMethodHS256)

	jwtToken.Claims = jwtgo.MapClaims{
		"username": u.Username,
		"roleId":   u.RoleID,
		"slug":     u.Slug,
		"orig_iat": time.Now().Add(time.Hour * 24 * 1).Unix(),
		"exp":      time.Now().Add(time.Hour * 24 * 1).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwtToken.SignedString([]byte("secret key"))
	return token
}

func (user *User) GetUserStatusAsString() string {
	switch user.Status {
	case "A":
		return "Active"
	case "I":
		return "Inctive"
	default:
		return "Unknown"
	}
}

func (user *User) IsAdmin() bool {

	if user.Role.Name == "admin" {
		return true
	}

	return false
}

func (user *User) IsStaff() bool {

	if user.Role.Name == "staff" {
		return true
	}

	return false
}

func (user *User) IsNotAdmin() bool {
	return !user.IsAdmin()
}

type SwagUser struct {
	SwagID
	SwagUserBody
}

type SwagUserPassword struct {
	Password string `json:"password" example:"pass1234"` // รหัสผ่าน
}

type SwagUserBody struct {
	Username  string `json:"username" example:"user01"`        // Username
	FirstName string `json:"firstName" example:"john"`         // ชื่อ
	LastName  string `json:"lastName" example:"doe"`           // นามสกุล
	Email     string `json:"email" example:"user01@email.com"` // อีเมล์
	Slug      string `json:"slug" example:"user01"`            // Slug
	Avatar    string `json:"avatar" example:"user01.png"`      // รูป Avatar
}

type SwagUserBodyIncludePassword struct {
	SwagUserBody
	SwagUserPassword
}

type SwagGetAllUsersResponse struct {
	SwagGetBase
	Data     []SwagUser   `json:"data"`
	PageMeta SwagPageMeta `json:"pageMeta"`
}

type SwagGetUserResponse struct {
	SwagGetBase
	Data SwagUser `json:"data"`
}

type SwagCreateUserResponse struct {
	SwagCreateBase
	Data SwagUser `json:"data"`
}

type SwagUpdateUserResponse struct {
	SwagUpdateBase
	Data SwagUser `json:"data"`
}

type SwagChangePasswordResponse struct {
	Success bool        `json:"success" example:"true"`                         // ผลการเรียกใช้งาน
	Message string      `json:"message" example:"Change Password Successfully"` // ข้อความตอบกลับ
	Data    interface{} `json:"data" `                                          // ข้อมูล
}

type SwagUploadAvatarResponse struct {
	Success bool        `json:"success" example:"true"`                         // ผลการเรียกใช้งาน
	Message string      `json:"message" example:"Uploaded Avatar Successfully"` // ข้อความตอบกลับ
	Data    interface{} `json:"data" `                                          // ข้อมูล
}
