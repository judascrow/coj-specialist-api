package models

import (
	"github.com/jinzhu/gorm"
)

type Status struct {
	gorm.Model
	Code string `json:"shortName" form:"shortName" gorm:"type:varchar(1);not null"`
	Name string `json:"status" form:"status" gorm:"type:varchar(50);not null"`
}
