package services

import (
	"github.com/judascrow/cojspcl-api/api/infrastructure"
	"github.com/judascrow/cojspcl-api/api/models"
)

func FindAllProvinces() ([]models.Province, error) {
	db := infrastructure.GetDB()
	var provinces []models.Province
	err := db.Find(&provinces).Error
	return provinces, err
}

func FindDistrictsByProvinceID(provinceID int) ([]models.District, error) {
	db := infrastructure.GetDB()
	var districts []models.District
	err := db.Where(models.District{ProvinceID: provinceID}).Find(&districts).Error
	return districts, err
}

func FindSubDistrictsByDistrictID(provinceID, districtID int) ([]models.SubDistrict, error) {
	db := infrastructure.GetDB()
	var subDistricts []models.SubDistrict
	err := db.Joins("left join districts on districts.id = sub_districts.district_id").Where("districts.province_id = ?", provinceID).Where(models.SubDistrict{DistrictID: districtID}).Find(&subDistricts).Error
	return subDistricts, err
}

func FindAllDistricts() ([]models.District, error) {
	db := infrastructure.GetDB()
	var districts []models.District
	err := db.Find(&districts).Error
	return districts, err
}

func FindAllSubDistricts() ([]models.SubDistrict, error) {
	db := infrastructure.GetDB()
	var subDistricts []models.SubDistrict
	err := db.Find(&subDistricts).Error
	return subDistricts, err
}
