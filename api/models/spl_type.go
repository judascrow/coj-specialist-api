package models

import (
	"github.com/jinzhu/gorm"
)

type SplType struct {
	gorm.Model
	NameTH string `json:"nameTH" form:"nameTH"`
	NameEN string `json:"nameEN" form:"nameEN"`
}

type SplSubType struct {
	gorm.Model
	NameTH    string `json:"nameTH" form:"nameTH"`
	NameEN    string `json:"nameEN" form:"nameEN"`
	SplTypeID int    `json:"splTypeID" form:"splTypeID"`

	SplType SplType `json:"splType" form:"splType"`
}

func (s SplType) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":     s.ID,
		"nameTH": s.NameTH,
		"nameEH": s.NameEN,
	}
}

func (s SplSubType) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":      s.ID,
		"nameTH":  s.NameTH,
		"nameEH":  s.NameEN,
		"splType": s.SplType.Serialize(),
	}
}
