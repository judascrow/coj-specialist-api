package services

import (
	"github.com/judascrow/cojspcl-api/api/infrastructure"
	"github.com/judascrow/cojspcl-api/api/models"
)

func FindAllSplTypes() ([]models.SplType, error) {
	db := infrastructure.GetDB()
	var splTypes []models.SplType
	err := db.Find(&splTypes).Error
	return splTypes, err
}

func FindAllSplSubTypes() ([]models.SplSubType, error) {
	db := infrastructure.GetDB()
	var splSubTypes []models.SplSubType
	err := db.Find(&splSubTypes).Error
	return splSubTypes, err
}

func FindAllSplSubTypesBySplTypeID(splTypeID int) ([]models.SplSubType, error) {
	db := infrastructure.GetDB()
	var splSubTypes []models.SplSubType
	err := db.Set("gorm:auto_preload", true).Where(models.SplSubType{SplTypeID: splTypeID}).Find(&splSubTypes).Error
	return splSubTypes, err
}

func GetSplTypeByID(id uint) (models.SplType, error) {
	db := infrastructure.GetDB()
	var splType models.SplType
	err := db.First(&splType, id).Error
	return splType, err
}

func GetSplSubTypeByID(id uint) (models.SplSubType, error) {
	db := infrastructure.GetDB()
	var splSubType models.SplSubType
	err := db.Set("gorm:auto_preload", true).First(&splSubType, id).Error
	return splSubType, err
}
