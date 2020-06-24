package controllers

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/judascrow/cojspcl-api/api/models"

	"github.com/judascrow/cojspcl-api/api/services"
	"github.com/judascrow/cojspcl-api/api/utils/messages"
	"github.com/judascrow/cojspcl-api/api/utils/responses"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary รายการผู้ใช้งาน
// @Description รายการผู้ใช้งาน
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SwagGetAllUsersResponse
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Security ApiKeyAuth
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	// Query Pages
	pageSizeStr := c.Query("pageSize")
	pageStr := c.Query("page")

	// Find Users
	users, pageMeta, err := services.FindAllUsers(pageSizeStr, pageStr)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
	}

	// Serialize
	length := len(users)
	UserSerialized := make([]map[string]interface{}, length, length)
	for i := 0; i < length; i++ {
		UserSerialized[i] = users[i].Serialize()
	}

	// Response
	responses.JSONLIST(c, http.StatusOK, "users", UserSerialized, messages.DataFound, pageMeta)
}

// @Summary ข้อมูลผู้ใช้งาน
// @Description ข้อมูลผู้ใช้งาน
// @Tags User
// @Accept  json
// @Produce  json
// @Param slug path string true "slug ผู้ใช้งาน"
// @Success 200 {object} models.SwagGetUserResponse
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Security ApiKeyAuth
// @Router /users/{slug} [get]
func GetUserBySlug(c *gin.Context) {
	// Get Slug from URI
	slug := c.Param("slug")

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

// @Summary เพิ่มผู้ใช้งาน
// @Description เพิ่มผู้ใช้งาน
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body models.SwagUserBodyIncludePassword true "เพิ่มผู้ใช้งาน"
// @Success 201 {object} models.SwagCreateUserResponse
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Security ApiKeyAuth
// @Router /users [post]
func CreateUser(c *gin.Context) {

	// Define struct user variable
	var user models.User

	// Map jsonBody to struct
	err := c.BindJSON(&user)
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, messages.ErrorsResponse(err))
		return
	}

	// Check username duplicate
	userCond := models.User{Username: user.Username}
	_, err = services.FindOneUser(userCond)
	if err == nil {
		errMessage := "ชื่อผู้ใช้งาน " + messages.IsAlreadyExists
		responses.ERROR(c, http.StatusBadRequest, errMessage)
		return
	}
	// Check email duplicate

	userCond = models.User{Email: user.Email}
	_, err = services.FindOneUser(userCond)
	if err == nil {
		errMessage := "email " + messages.IsAlreadyExists
		responses.ERROR(c, http.StatusBadRequest, errMessage)
		return
	}

	// Generate password
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

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

// @Summary แก้ไขผู้ใช้งาน
// @Description แก้ไขผู้ใช้งาน
// @Tags User
// @Accept  json
// @Produce  json
// @Param slug path string true "slug ผู้ใช้งาน"
// @Param user body models.SwagUserBody true "แก้ไขผู้ใช้งาน"
// @Success 201 {object} models.SwagUpdateUserResponse
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Security ApiKeyAuth
// @Router /users/{slug} [put]
func UpdateUser(c *gin.Context) {

	slug := c.Param("slug")

	if !ClaimsOwner(c, slug) {
		responses.ERROR(c, http.StatusForbidden, messages.NotPermission)
		return
	}

	var userData *models.User
	err := c.BindJSON(&userData)
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, messages.ErrorsResponse(err))
		return
	}

	var user models.User
	userData.Username = ""
	userData.Password = ""

	if user, err = services.UpdateUser(slug, userData); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	responses.JSON(c, http.StatusOK, user.Serialize(), messages.Updated+messages.User+messages.Success)

}

// @Summary ลบผู้ใช้งาน
// @Description ลบผู้ใช้งาน
// @Tags User
// @Accept  json
// @Produce  json
// @Param slug path string true "slug ผู้ใช้งาน"
// @Success 201 {object} models.SwagDeleteBase
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Security ApiKeyAuth
// @Router /users/{slug} [delete]
func DeleteUser(c *gin.Context) {
	slug := c.Param("slug")
	err := services.DeleteUser(&models.User{Slug: slug})
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}
	responses.JSONNODATA(c, http.StatusNoContent, messages.Deleted+messages.User+messages.Success)
}

// @Summary เปลี่ยนรหัสผ่าน
// @Description เปลี่ยนรหัสผ่าน
// @Tags User
// @Accept  json
// @Produce  json
// @Param slug path string true "slug ผู้ใช้งาน"
// @Param user body models.ChangePassword true "เปลี่ยนรหัสผ่าน"
// @Success 201 {object} models.SwagChangePasswordResponse
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Security ApiKeyAuth
// @Router /users/{slug}/password [post]
func ChangePassword(c *gin.Context) {

	slug := c.Param("slug")

	var requestBody models.ChangePassword
	err := c.BindJSON(&requestBody)
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, messages.ErrorsResponse(err))
		return
	}

	user, err := services.FindOneUserBySlug(slug)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	if !ClaimsOwner(c, slug) {
		responses.ERROR(c, http.StatusForbidden, messages.NotPermission)
		return
	}

	password := user.Password
	if requestBody.CurrentPassword == "" && requestBody.NewPassword != "" {
		err = errors.New("Please Provide current_password")
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}
	if requestBody.CurrentPassword != "" && requestBody.NewPassword == "" {
		err = errors.New("Please Provide new_password")
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}
	if requestBody.CurrentPassword != "" && requestBody.NewPassword != "" {
		//Also check if the new password
		if len(requestBody.NewPassword) < 6 {
			err = errors.New("Password should be atleast 6 characters")
			responses.ERROR(c, http.StatusBadRequest, err.Error())
			return
		}
		//if they do, check that the former password is correct
		err = verifyPassword(user.Password, requestBody.CurrentPassword)
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			err = errors.New("The password not correct")
			responses.ERROR(c, http.StatusBadRequest, err.Error())
			return
		}

		// Generate password
		bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			responses.ERROR(c, http.StatusBadRequest, err.Error())
			return
		}
		password = string(bcryptPassword)
	}

	userData := models.User{
		Password: password,
	}

	if user, err = services.UpdateUser(slug, userData); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}
	responses.JSONNODATA(c, http.StatusOK, messages.ChangePasswordSuccess)

}

// @Summary อัพโหลดรูป avatar
// @Description อัพโหลดรูป avatar
// @Tags User
// @Accept  json
// @Produce  json
// @Param slug path string true "slug ผู้ใช้งาน"
// @Param user body models.UploadAvatar true "อัพโหลดรูป avatar"
// @Success 201 {object} models.SwagUploadAvatarResponse
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Security ApiKeyAuth
// @Router /users/{slug}/avatar [post]
func UploadAvatar(c *gin.Context) {

	slug := c.Param("slug")

	if !ClaimsOwner(c, slug) {
		responses.ERROR(c, http.StatusForbidden, messages.NotPermission)
		return
	}

	user, err := services.FindOneUserBySlug(slug)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	avatar, err := c.FormFile("avatar")
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	if avatar != nil {
		imgNameAvatar := randomString(16) + ".png"
		dirPath := filepath.Join(".", "upload", "avatar")
		filePathAvatar := filepath.Join(dirPath, imgNameAvatar)
		if _, err = os.Stat(dirPath); os.IsNotExist(err) {
			err = os.MkdirAll(dirPath, os.ModeDir)
			if err != nil {
				responses.ERROR(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
		if err := c.SaveUploadedFile(avatar, filePathAvatar); err != nil {
			responses.ERROR(c, http.StatusBadRequest, err.Error())
			return
		}

		user.Avatar = string(filepath.Separator) + filePathAvatar
	}

	userData := models.User{
		Avatar: user.Avatar,
	}

	if user, err = services.UpdateUser(slug, userData); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}
	responses.JSON(c, http.StatusOK, user.Serialize(), messages.UploadedAvatarSuccess)

}
