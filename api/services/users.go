package services

import (
	"github.com/judascrow/cojspcl-api/api/infrastructure"
	"github.com/judascrow/cojspcl-api/api/models"
)

func FindAllUsers(pageSizeStr, pageStr string) ([]models.User, PageMeta, error) {
	db := infrastructure.GetDB()
	pageMeta := getPageMeta(pageSizeStr, pageStr)

	var users []models.User
	var count int

	db.Model(&models.User{}).Count(&count)
	err := db.Set("gorm:auto_preload", true).Offset((pageMeta.Page - 1) * pageMeta.PageSize).Limit(pageMeta.PageSize).Find(&users).Error

	pageMeta.TotalItemsCount = count
	pageMeta.CurrentItemsCount = len(users)

	return users, pageMeta, err
}

func FindOneUserBySlug(slug string) (models.User, error) {
	db := infrastructure.GetDB()
	var user models.User
	err := db.Set("gorm:auto_preload", true).Where(&models.User{Slug: slug}).First(&user).Error
	return user, err
}

func FindOneUser(condition interface{}) (models.User, error) {
	db := infrastructure.GetDB()
	var user models.User
	err := db.Set("gorm:auto_preload", true).Where(condition).First(&user).Error
	return user, err
}

func UpdateUser(slug string, data interface{}) (models.User, error) {
	db := infrastructure.GetDB()
	var user models.User
	err := db.Model(user).Set("gorm:auto_preload", true).Where(models.User{Slug: slug}).Update(data).Take(&user).Error
	return user, err
}

func DeleteUser(condition interface{}) error {
	db := infrastructure.GetDB()
	err := db.Unscoped().Where(condition).Delete(models.User{}).Error
	return err
}
