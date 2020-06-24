package controllers

import (
	"net/http"

	"github.com/judascrow/cojspcl-api/api/services"
	"github.com/judascrow/cojspcl-api/api/utils/messages"
	"github.com/judascrow/cojspcl-api/api/utils/responses"

	"github.com/gin-gonic/gin"
)

// @Summary รายการสิทธิ์การใช้งาน
// @Description รายการสิทธิ์การใช้งาน
// @Tags Role
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SwagGetRoleResponse
// @Failure 400 {object} models.SwagError400
// @Failure 404 {object} models.SwagError404
// @Failure 500 {object} models.SwagError500
// @Security ApiKeyAuth
// @Router /roles [get]
func GetAllRoles(c *gin.Context) {
	// Find Users
	roles, err := services.FindAllRoles()
	if err != nil {
		responses.ERROR(c, http.StatusNotFound, messages.NotFound)
	}

	// Serialize
	length := len(roles)
	RolesSerialized := make([]map[string]interface{}, length, length)
	for i := 0; i < length; i++ {
		RolesSerialized[i] = roles[i].Serialize()
	}

	// Response
	responses.JSON(c, http.StatusOK, RolesSerialized, messages.DataFound)
}
