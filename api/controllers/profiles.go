package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/judascrow/cojspcl-api/api/models"
	"github.com/judascrow/cojspcl-api/api/services"
	"github.com/judascrow/cojspcl-api/api/utils/messages"
	"github.com/judascrow/cojspcl-api/api/utils/responses"
	jwt "github.com/judascrow/gomiddlewares/jwt"

	"github.com/gin-gonic/gin"
)

func GetAllProfileLists(c *gin.Context) {

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

	// Serialize
	length := len(reqforms)
	reqformsSerialized := make([]map[string]interface{}, length, length)
	for i := 0; i < length; i++ {
		reqformsSerialized[i] = reqforms[i].SerializeList()
	}

	responses.JSON(c, http.StatusOK, reqformsSerialized, messages.DataFound)
}

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

func GetProfile(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	profile, err := services.GetProfileByID(uint(id))
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	replaceAllFlag := -1
	profile.FileAttachIdcardURL = strings.Replace(profile.FileAttachIdcardURL, "\\", "/", replaceAllFlag)

	responses.JSON(c, http.StatusOK, profile, messages.DataFound)
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

func UpdateProfile(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
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

	fileAttachIdcard, err := UploadFilePDF(c, profileRequest.FileAttachIdcard, uint(id))
	fileAttachHouse, err := UploadFilePDF(c, profileRequest.FileAttachHouse, uint(id))
	fileAttachGovCard, err := UploadFilePDF(c, profileRequest.FileAttachGovCard, uint(id))
	fileAttachQualification, err := UploadFilePDF(c, profileRequest.FileAttachQualification, uint(id))

	var profileData models.Profile
	// Map jsonBody to struct
	err = c.ShouldBind(&profileData)
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, messages.ErrorsResponse(err))
		return
	}
	profileData.UserId = int(uint(id))
	profileData.FileAttachIdcardURL = fileAttachIdcard
	profileData.FileAttachHouseURL = fileAttachHouse
	profileData.FileAttachGovCardURL = fileAttachGovCard
	profileData.FileAttachQualificationURL = fileAttachQualification

	if profileData.StatusReqform == "approved" {
		t := true
		profileData.IsSpecialist = &t
	} else {
		t := false
		profileData.IsSpecialist = &t
	}

	profileData.UserId = 0

	if profileData, err = services.UpdateProfile(uint(id), profileData); err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	// Find User
	profile, err := services.GetProfileByID(profileData.ID)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	responses.JSON(c, http.StatusOK, profile, messages.Updated+"ข้อมูล"+messages.Success)
}

func DeleteProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	err = services.DeleteProfile(uint(id))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}
	responses.JSONNODATA(c, http.StatusNoContent, messages.Deleted+"ข้อมูล"+messages.Success)
}
