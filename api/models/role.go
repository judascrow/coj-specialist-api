package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique" example:"ROLE_USER"`
	Description string `json:"description" example:"Only for standard users"`
}

func (r Role) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":          r.ID,
		"name":        r.Name,
		"description": r.Description,
	}
}

type SwagGetRoleResponse struct {
	SwagGetBase
	Data Role `json:"data"`
}
