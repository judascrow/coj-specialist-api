package controllers

import (
	"net/http"
	"strconv"

	"github.com/judascrow/cojspcl-api/api/services"
	"github.com/judascrow/cojspcl-api/api/utils/messages"
	"github.com/judascrow/cojspcl-api/api/utils/responses"

	"github.com/gin-gonic/gin"
)

func GetAllSplTypes(c *gin.Context) {
	splTypes, err := services.FindAllSplTypes()
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
	}

	// Serialize
	length := len(splTypes)
	splTypesSerialized := make([]map[string]interface{}, length, length)
	for i := 0; i < length; i++ {
		splTypesSerialized[i] = splTypes[i].Serialize()
	}

	// Response
	responses.JSON(c, http.StatusOK, splTypesSerialized, messages.DataFound)
}

func GetAllSplSubTypes(c *gin.Context) {
	splSubTypes, err := services.FindAllSplSubTypes()
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
	}

	// Serialize
	length := len(splSubTypes)
	splSubTypesSerialized := make([]map[string]interface{}, length, length)
	for i := 0; i < length; i++ {
		splSubTypesSerialized[i] = splSubTypes[i].Serialize()
	}

	// Response
	responses.JSON(c, http.StatusOK, splSubTypesSerialized, messages.DataFound)
}

func GetSplSubTypesBySplTypeID(c *gin.Context) {
	splTypeID, err := strconv.Atoi(c.Param("splTypeID"))

	if err != nil {
		responses.ERROR(c, http.StatusBadRequest, err.Error())
		return
	}

	splSubTypes, err := services.FindAllSplSubTypesBySplTypeID(splTypeID)
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
	}
	// Serialize
	length := len(splSubTypes)
	splSubTypesSerialized := make([]map[string]interface{}, length, length)
	for i := 0; i < length; i++ {
		splSubTypesSerialized[i] = splSubTypes[i].Serialize()
	}

	// Response
	responses.JSON(c, http.StatusOK, splSubTypesSerialized, messages.DataFound)
}
