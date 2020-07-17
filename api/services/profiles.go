package services

import (
	"github.com/judascrow/cojspcl-api/api/infrastructure"
	"github.com/judascrow/cojspcl-api/api/models"
)

func GetAllProfiles() ([]models.Profile, error) {
	db := infrastructure.GetDB()
	var profiles []models.Profile
	err := db.Find(&profiles).Error
	return profiles, err
}

func GetProfileByID(id uint) (models.Profile, error) {
	db := infrastructure.GetDB()
	var profile models.Profile
	err := db.First(&profile, id).Error
	return profile, err
}

func GetProfileCondition(Condition interface{}) ([]models.Profile, error) {
	db := infrastructure.GetDB()
	var profiles []models.Profile
	err := db.Where(Condition).Find(&profiles).Error
	return profiles, err
}
