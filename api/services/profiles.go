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
	err := db.Preload("SplSubType").Preload("SplSubType.SplType").Where(Condition).Find(&profiles).Error
	return profiles, err
}

func UpdateProfile(id uint, data interface{}) (models.Profile, error) {
	db := infrastructure.GetDB()
	var profile models.Profile
	err := db.Model(profile).Set("gorm:auto_preload", true).Where("id = ?", id).Update(data).Take(&profile).Error
	return profile, err
}

func DeleteProfile(id uint) error {
	db := infrastructure.GetDB()
	err := db.Unscoped().Where("id = ?", id).Delete(models.Profile{}).Error
	return err
}
