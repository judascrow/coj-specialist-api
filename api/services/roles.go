package services

import (
	"github.com/judascrow/cojspcl-api/api/infrastructure"
	"github.com/judascrow/cojspcl-api/api/models"
)

func FindAllRoles() ([]models.Role, error) {
	db := infrastructure.GetDB()

	var roles []models.Role

	err := db.Find(&roles).Error

	return roles, err
}
