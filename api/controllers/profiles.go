package controllers

import (
	"fmt"
	"net/http"

	"github.com/judascrow/cojspcl-api/api/models"
	"github.com/judascrow/cojspcl-api/api/services"
	"github.com/judascrow/cojspcl-api/api/utils/messages"
	"github.com/judascrow/cojspcl-api/api/utils/responses"
	jwt "github.com/judascrow/gomiddlewares/jwt"

	"github.com/gin-gonic/gin"
)

func GetAllReqforms(c *gin.Context) {

	var condition models.Profile
	if err := c.BindQuery(&condition); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	reqforms, err := services.GetProfileCondition(condition)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}
	responses.JSON(c, http.StatusOK, reqforms, messages.DataFound)
}

func CreateProfile(c *gin.Context) {

	claims := jwt.ExtractClaims(c)
	slug := claims["slug"].(string)

	user, err := services.FindOneUserBySlug(slug)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	// Check username duplicate
	profileCond := models.Profile{UserId: int(user.ID)}
	profilesCheck, err := services.GetProfileCondition(profileCond)
	if err == nil && len(profilesCheck) > 0 {
		errMessage := "ท่านเคยส่งข้อมูลคำขอขึ้นทะเบียนแล้ว"
		responses.ERROR(c, http.StatusBadRequest, errMessage)
		return
	}

	var profileRequest models.ProfileRequest

	// Map jsonBody to struct
	err = c.ShouldBind(&profileRequest)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	fileAttachIdcard, err := UploadFilePDF(c, profileRequest.FileAttachIdcard, user.ID)
	fileAttachHouse, err := UploadFilePDF(c, profileRequest.FileAttachHouse, user.ID)
	fileAttachGovCard, err := UploadFilePDF(c, profileRequest.FileAttachGovCard, user.ID)
	fileAttachQualification, err := UploadFilePDF(c, profileRequest.FileAttachQualification, user.ID)

	var profileData models.Profile
	// Map jsonBody to struct
	err = c.ShouldBind(&profileData)
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, messages.ErrorsResponse(err))
		return
	}
	profileData.UserId = int(user.ID)
	profileData.FileAttachIdcardURL = fileAttachIdcard
	profileData.FileAttachHouseURL = fileAttachHouse
	profileData.FileAttachGovCardURL = fileAttachGovCard
	profileData.FileAttachQualificationURL = fileAttachQualification

	// Create user
	if err = services.CreateOne(&profileData); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	// Find User
	profile, err := services.GetProfileByID(profileData.ID)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	// Response
	responses.JSON(c, http.StatusCreated, profile, "ส่งคำขอขึ้นทะเบียนเรียบร้อยแล้ว")
}
