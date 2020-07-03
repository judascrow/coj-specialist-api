package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	UserId              int        `json:"userId" form:"userId"`
	Status              string     `json:"status" form:"status" sql:"type:enum('A','I');DEFAULT:'A'"`
	IsSpecialist        *bool      `json:"isSpecialist" form:"isSpecialist"`
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
	AddressNo           string     `json:"addressNo" form:"addressNo"`
	AddressMoo          string     `json:"addressMoo" form:"addressMoo"`
	AddressSoi          string     `json:"addressSoi" form:"addressSoi"`
	AddressRoad         string     `json:"addressRoad" form:"addressRoad"`
	AddressSubDistrict  string     `json:"addressSubDistrict" form:"addressSubDistrict"`
	AddressDistrict     string     `json:"addressDistrict" form:"addressDistrict"`
	AddressProvince     string     `json:"addressProvince" form:"addressProvince"`
	AddressZipcode      string     `json:"addressZipcode" form:"addressZipcode"`
	AddressTel          string     `json:"addressTel" form:"addressTel"`
	AddressFax          string     `json:"addressFax" form:"addressFax"`
	AddressEmail        string     `json:"addressEmail" form:"addressEmail"`
	ContactNo           string     `json:"contactNo" form:"contactNo"`
	ContactMoo          string     `json:"contactMoo" form:"contactMoo"`
	ContactSoi          string     `json:"contactSoi" form:"contactSoi"`
	ContactRoad         string     `json:"contactRoad" form:"contactRoad"`
	ContactSubDistrict  string     `json:"contactSubDistrict" form:"contactSubDistrict"`
	ContactDistrict     string     `json:"contactDistrict" form:"contactDistrict"`
	ContactProvince     string     `json:"contactProvince" form:"contactProvince"`
	ContactZipcode      string     `json:"contactZipcode" form:"contactZipcode"`
	ContactTel          string     `json:"contactTel" form:"contactTel"`
	ContactFax          string     `json:"contactFax" form:"contactFax"`
	ContactEmail        string     `json:"contactEmail" form:"contactEmail"`
	WorkOccupation      string     `json:"workOccupation" form:"workOccupation"`
	WorkPosition        string     `json:"workPosition" form:"workPosition"`
	WorkPlaces          string     `json:"workPlaces" form:"workPlaces"`
	WorkRoad            string     `json:"workRoad" form:"workRoad"`
	WorkSubDistrict     string     `json:"workSubDistrict" form:"workSubDistrict"`
	WorkDistrict        string     `json:"workDistrict" form:"workDistrict"`
	WorkProvince        string     `json:"workProvince" form:"workProvince"`
	WorkZipcode         string     `json:"workZipcode" form:"workZipcode"`
	WorkTel             string     `json:"workTel" form:"workTel"`
	WorkFax             string     `json:"workFax" form:"workFax"`
	WorkEmail           string     `json:"workEmail" form:"workEmail"`
	BossFirstName       string     `json:"bossFirstName" form:"bossFirstName"`
	BossLastName        string     `json:"bossLastName" form:"bossLastName"`
	BossNo              string     `json:"bossNo" form:"bossNo"`
	BossMoo             string     `json:"bossMoo" form:"bossMoo"`
	BossSoi             string     `json:"bossSoi" form:"bossSoi"`
	BossRoad            string     `json:"bossRoad" form:"bossRoad"`
	BossSubDistrict     string     `json:"bossSubDistrict" form:"bossSubDistrict"`
	BossDistrict        string     `json:"bossDistrict" form:"bossDistrict"`
	BossProvince        string     `json:"bossProvince" form:"bossProvince"`
	BossZipcode         string     `json:"bossZipcode" form:"bossZipcode"`
	BossTel             string     `json:"bossTel" form:"bossTel"`
	BossFax             string     `json:"bossFax" form:"bossFax"`
	BossEmail           string     `json:"bossEmail" form:"bossEmail"`
	WorkExperience      string     `json:"workExperience" form:"workExperience"`
	RegisWork           string     `json:"regisWork" form:"regisWork"`
	RegisQualification  string     `json:"regisQualification" form:"regisQualification"`
	RegisDocument       string     `json:"regisDocument" form:"regisDocument"`
	RegisEver           string     `json:"regisEver" form:"regisEver"`
	RegisEverYear       string     `json:"regisEverYear" form:"regisEverYear"`
	RegisEverPass       string     `json:"regisEverPass" form:"regisEverPass"`
	RegisEverPassNo     string     `json:"regisEverPassNo" form:"regisEverPassNo"`
	RegisEverNopass     string     `json:"regisEverNopass" form:"regisEverNopass"`
	RegisEverNopassDesc string     `json:"regisEverNopassDesc" form:"regisEverNopassDesc"`
	FileAttach1         string     `json:"fileAttach1" form:"fileAttach1"`
	Remark              string     `json:"remark" form:"remark"`
}
