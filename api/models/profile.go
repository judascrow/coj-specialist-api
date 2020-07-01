package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	UserId              int        `json:"userId" form:"userId"`
	Idcard              string     `json:"idcard" form:"idcard"`
	GovCard             string     `json:"govCard" form:"govCard"`
	CardExpire          *time.Time `json:"cardExpire" form:"cardExpire"`
	PrefixName          string     `json:"prefixName" form:"prefixName"`
	FirstName           string     `json:"firstName" form:"firstName"`
	LastName            string     `json:"lastName" form:"lastName"`
	BirthDate           *time.Time `json:"birthDate" form:"birthDate"`
	Race                string     `json:"race" form:"race"`
	Nation              string     `json:"nation" form:"nation"`
	DomicileNo          string     `json:"domicileNo" form:"domicileNo"`
	DomicileMoo         string     `json:"domicileMoo" form:"domicileMoo"`
	DomicileSoi         string     `json:"domicileSoi" form:"domicileSoi"`
	DomicileRoad        string     `json:"domicileRoad" form:"domicileRoad"`
	DomicileProvince    string     `json:"domicileProvince" form:"domicileProvince"`
	DomicileDistrict    string     `json:"domicileDistrict" form:"domicileDistrict"`
	DomicileSubDistrict string     `json:"domicileSubDistrict" form:"domicileSubDistrict"`
	DomicileZipcode     string     `json:"domicileZipcode" form:"domicileZipcode"`
	DomicileTel         string     `json:"domicileTel" form:"domicileTel"`
	DomicileFax         string     `json:"domicileFax" form:"domicileFax"`
	DomicileEmail       string     `json:"domicileEmail" form:"domicileEmail"`
}
