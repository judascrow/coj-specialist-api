package models

import (
	"github.com/jinzhu/gorm"
)

type Province struct {
	gorm.Model
	Code   int    `json:"code" form:"code"`
	NameTH string `json:"nameTH" form:"nameTH"`
	NameEN string `json:"nameEN" form:"nameEN"`
}

type District struct {
	gorm.Model
	Code       int    `json:"code" form:"code"`
	NameTH     string `json:"nameTH" form:"nameTH"`
	NameEN     string `json:"nameEN" form:"nameEN"`
	ProvinceID int    `json:"provinceID" form:"provinceID"`
}

type SubDistrict struct {
	gorm.Model
	Code       int    `json:"code" form:"code"`
	NameTH     string `json:"nameTH" form:"nameTH"`
	NameEN     string `json:"nameEN" form:"nameEN"`
	ZipCode    int    `json:"zipCode" form:"zipCode"`
	DistrictID int    `json:"districtID" form:"districtID"`
}

func (p Province) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":     p.ID,
		"nameTH": p.NameTH,
	}
}

func (d District) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":     d.ID,
		"nameTH": d.NameTH,
	}
}

func (s SubDistrict) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":     s.ID,
		"nameTH": s.NameTH,
	}
}
