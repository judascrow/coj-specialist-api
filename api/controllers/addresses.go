package controllers

import (
	"net/http"
	"strconv"

	"github.com/judascrow/cojspcl-api/api/services"
	"github.com/judascrow/cojspcl-api/api/utils/messages"
	"github.com/judascrow/cojspcl-api/api/utils/responses"

	"github.com/gin-gonic/gin"
)

func GetAllProvinces(c *gin.Context) {
	// Find Provinces
	provinces, err := services.FindAllProvinces()
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
	}

	// Serialize
	length := len(provinces)
	provincesSerialized := make([]map[string]interface{}, length, length)
	for i := 0; i < length; i++ {
		provincesSerialized[i] = provinces[i].Serialize()
	}

	// Response
	responses.JSON(c, http.StatusOK, provincesSerialized, messages.DataFound)
}

func GetProvinceByID(c *gin.Context) {
	provinceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	province, err := services.GetProvinceByID(uint(provinceID))
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	responses.JSON(c, http.StatusOK, province.Serialize(), messages.DataFound)
}

func GetDistrictsByProvinceID(c *gin.Context) {
	provinceID, err := strconv.Atoi(c.Param("provinceID"))

	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	districts, err := services.FindDistrictsByProvinceID(provinceID)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
	}
	// Serialize
	length := len(districts)
	districtsSerialized := make([]map[string]interface{}, length, length)
	for i := 0; i < length; i++ {
		districtsSerialized[i] = districts[i].Serialize()
	}

	// Response
	responses.JSON(c, http.StatusOK, districtsSerialized, messages.DataFound)
}

func GetSubDistrictsByDistrictID(c *gin.Context) {
	provinceID, err := strconv.Atoi(c.Param("provinceID"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	districtID, err := strconv.Atoi(c.Param("districtID"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	subDistricts, err := services.FindSubDistrictsByDistrictID(provinceID, districtID)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
	}
	// Serialize
	length := len(subDistricts)
	subDistrictsSerialized := make([]map[string]interface{}, length, length)
	for i := 0; i < length; i++ {
		subDistrictsSerialized[i] = subDistricts[i].Serialize()
	}

	// Response
	responses.JSON(c, http.StatusOK, subDistrictsSerialized, messages.DataFound)
}

func GetDistrictByID(c *gin.Context) {
	districtID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	district, err := services.GetDistrictByID(uint(districtID))
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	responses.JSON(c, http.StatusOK, district.Serialize(), messages.DataFound)
}

func GetSubDistrictByID(c *gin.Context) {
	subDistrictID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	subDistrict, err := services.GetSubDistrictByID(uint(subDistrictID))
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
		return
	}

	responses.JSON(c, http.StatusOK, subDistrict.Serialize(), messages.DataFound)
}
