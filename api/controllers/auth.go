package controllers

import (
	"net/http"

	jwt "github.com/judascrow/gomiddlewares/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/judascrow/cojspcl-api/api/models"
	"github.com/judascrow/cojspcl-api/api/services"
	"github.com/judascrow/cojspcl-api/api/utils/messages"
	"github.com/judascrow/cojspcl-api/api/utils/responses"

	"github.com/gin-gonic/gin"
)

// @Summary ข้อมูลตนเอง
// @Description ข้อมูลตนเอง
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SwagGetUserResponse
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Security ApiKeyAuth
// @Router /auth/me [get]
func GetUserMe(c *gin.Context) {

	claims := jwt.ExtractClaims(c)

	slug := claims["slug"].(string)

	if !ClaimsOwner(c, slug) {
		responses.ERROR(c, http.StatusForbidden, messages.NotPermission)
		return
	}

	// Find User
	user, err := services.FindOneUserBySlug(slug)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	// Response
	responses.JSON(c, http.StatusOK, user.Serialize(), messages.DataFound)
}

// @Summary สมัครเข้าใช้งาน
// @Description สมัครเข้าใช้งาน
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param user body models.RegisterUser true "สมัครเข้าใช้งาน"
// @Success 201 {object} models.SwagCreateUserResponse
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Router /auth/register [post]
func Register(c *gin.Context) {
	// Define struct user variable
	var register models.RegisterUser

	// Map jsonBody to struct
	err := c.BindJSON(&register)
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, messages.ErrorsResponse(err))
		return
	}

	// Check username duplicate
	userCond := models.User{Username: register.Username}
	_, err = services.FindOneUser(userCond)
	if err == nil {
		errMessage := "ชื่อผู้ใช้งาน " + messages.IsAlreadyExists
		responses.ERROR(c, http.StatusBadRequest, errMessage)
		return
	}
	// Check email duplicate

	userCond = models.User{Email: register.Email}
	_, err = services.FindOneUser(userCond)
	if err == nil {
		errMessage := "email " + messages.IsAlreadyExists
		responses.ERROR(c, http.StatusBadRequest, errMessage)
		return
	}

	// Generate password
	password, _ := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)

	user := models.User{
		Username:  register.Username,
		Password:  string(password),
		Email:     register.Email,
		FirstName: register.FirstName,
		LastName:  register.LastName,
		RoleID:    3,
	}

	// Create user
	if err = services.CreateOne(&user); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	// Find User
	user, err = services.FindOneUserBySlug(user.Slug)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	// Response
	responses.JSON(c, http.StatusCreated, user.Serialize(), messages.Created+messages.User+messages.Success)
}
