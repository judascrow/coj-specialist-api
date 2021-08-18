package models

import (
	"mime/multipart"
	"time"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	UserId                     int        `json:"userId" form:"userId"  gorm:"unique_index;DEFAULT:null"`
	Status                     string     `json:"status" form:"status" sql:"type:enum('A','I');DEFAULT:'A'"`
	IsSpecialist               *bool      `json:"isSpecialist" form:"isSpecialist" gorm:"DEFAULT:false"`
	StatusReqform              string     `json:"statusReqform" form:"statusReqform" sql:"type:enum('edit','checking','verify','approved');DEFAULT:'checking'"`
	IdCard                     string     `json:"idCard" form:"idCard"`
	GovCard                    string     `json:"govCard" form:"govCard"`
	CardExpire                 *time.Time `json:"cardExpire" form:"cardExpire" gorm:"type:date"`
	PrefixName                 string     `json:"prefixName" form:"prefixName"`
	FirstName                  string     `json:"firstName" form:"firstName"`
	LastName                   string     `json:"lastName" form:"lastName"`
	BirthDate                  *time.Time `json:"birthDate" form:"birthDate" gorm:"type:date"`
	Race                       string     `json:"race" form:"race"`
	Nation                     string     `json:"nation" form:"nation"`
	DomicileNo                 string     `json:"domicileNo" form:"domicileNo"`
	DomicileMoo                string     `json:"domicileMoo" form:"domicileMoo"`
	DomicileSoi                string     `json:"domicileSoi" form:"domicileSoi"`
	DomicileRoad               string     `json:"domicileRoad" form:"domicileRoad"`
	DomicileProvince           int        `json:"domicileProvince" form:"domicileProvince"`
	DomicileDistrict           int        `json:"domicileDistrict" form:"domicileDistrict"`
	DomicileSubDistrict        int        `json:"domicileSubDistrict" form:"domicileSubDistrict"`
	DomicileZipcode            string     `json:"domicileZipcode" form:"domicileZipcode"`
	DomicileTel                string     `json:"domicileTel" form:"domicileTel"`
	DomicileFax                string     `json:"domicileFax" form:"domicileFax"`
	DomicileEmail              string     `json:"domicileEmail" form:"domicileEmail"`
	AddressNo                  string     `json:"addressNo" form:"addressNo"`
	AddressMoo                 string     `json:"addressMoo" form:"addressMoo"`
	AddressSoi                 string     `json:"addressSoi" form:"addressSoi"`
	AddressRoad                string     `json:"addressRoad" form:"addressRoad"`
	AddressSubDistrict         int        `json:"addressSubDistrict" form:"addressSubDistrict"`
	AddressDistrict            int        `json:"addressDistrict" form:"addressDistrict"`
	AddressProvince            int        `json:"addressProvince" form:"addressProvince"`
	AddressZipcode             string     `json:"addressZipcode" form:"addressZipcode"`
	AddressTel                 string     `json:"addressTel" form:"addressTel"`
	AddressFax                 string     `json:"addressFax" form:"addressFax"`
	AddressEmail               string     `json:"addressEmail" form:"addressEmail"`
	ContactNo                  string     `json:"contactNo" form:"contactNo"`
	ContactMoo                 string     `json:"contactMoo" form:"contactMoo"`
	ContactSoi                 string     `json:"contactSoi" form:"contactSoi"`
	ContactRoad                string     `json:"contactRoad" form:"contactRoad"`
	ContactSubDistrict         int        `json:"contactSubDistrict" form:"contactSubDistrict"`
	ContactDistrict            int        `json:"contactDistrict" form:"contactDistrict"`
	ContactProvince            int        `json:"contactProvince" form:"contactProvince"`
	ContactZipcode             string     `json:"contactZipcode" form:"contactZipcode"`
	ContactTel                 string     `json:"contactTel" form:"contactTel"`
	ContactFax                 string     `json:"contactFax" form:"contactFax"`
	ContactEmail               string     `json:"contactEmail" form:"contactEmail"`
	WorkOccupation             string     `json:"workOccupation" form:"workOccupation"`
	WorkPosition               string     `json:"workPosition" form:"workPosition"`
	WorkPlaces                 string     `json:"workPlaces" form:"workPlaces"`
	WorkRoad                   string     `json:"workRoad" form:"workRoad"`
	WorkSubDistrict            int        `json:"workSubDistrict" form:"workSubDistrict"`
	WorkDistrict               int        `json:"workDistrict" form:"workDistrict"`
	WorkProvince               int        `json:"workProvince" form:"workProvince"`
	WorkZipcode                string     `json:"workZipcode" form:"workZipcode"`
	WorkTel                    string     `json:"workTel" form:"workTel"`
	WorkFax                    string     `json:"workFax" form:"workFax"`
	WorkEmail                  string     `json:"workEmail" form:"workEmail"`
	BossFirstName              string     `json:"bossFirstName" form:"bossFirstName"`
	BossLastName               string     `json:"bossLastName" form:"bossLastName"`
	BossNo                     string     `json:"bossNo" form:"bossNo"`
	BossMoo                    string     `json:"bossMoo" form:"bossMoo"`
	BossSoi                    string     `json:"bossSoi" form:"bossSoi"`
	BossRoad                   string     `json:"bossRoad" form:"bossRoad"`
	BossSubDistrict            int        `json:"bossSubDistrict" form:"bossSubDistrict"`
	BossDistrict               int        `json:"bossDistrict" form:"bossDistrict"`
	BossProvince               int        `json:"bossProvince" form:"bossProvince"`
	BossZipcode                string     `json:"bossZipcode" form:"bossZipcode"`
	BossTel                    string     `json:"bossTel" form:"bossTel"`
	BossFax                    string     `json:"bossFax" form:"bossFax"`
	BossEmail                  string     `json:"bossEmail" form:"bossEmail"`
	WorkExperience             string     `json:"workExperience" form:"workExperience"`
	SplTypeID                  int        `json:"splTypeID" form:"splTypeID"`
	SplSubTypeID               int        `json:"splSubTypeID" form:"splSubTypeID"`
	RegisQualification         string     `json:"regisQualification" form:"regisQualification"`
	RegisDocument              string     `json:"regisDocument" form:"regisDocument"`
	RegisEver                  string     `json:"regisEver" form:"regisEver"`
	RegisEverYear              string     `json:"regisEverYear" form:"regisEverYear"`
	RegisEverPass              string     `json:"regisEverPass" form:"regisEverPass"`
	RegisEverPassNo            string     `json:"regisEverPassNo" form:"regisEverPassNo"`
	RegisEverNopass            string     `json:"regisEverNopass" form:"regisEverNopass"`
	RegisEverNopassDesc        string     `json:"regisEverNopassDesc" form:"regisEverNopassDesc"`
	FileAttachIdcardURL        string     `json:"fileAttachIdcardURL" form:"fileAttachIdcardURL"`
	FileAttachHouseURL         string     `json:"fileAttachHouseURL" form:"fileAttachHouseURL"`
	FileAttachGovCardURL       string     `json:"fileAttachGovCardURL" form:"fileAttachGovCardURL"`
	FileAttachQualificationURL string     `json:"fileAttachQualificationURL" form:"fileAttachQualificationURL"`
	Remark                     string     `json:"remark" form:"remark"`

	SplSubType SplSubType `json:"SplSubType" form:"SplSubType"`
}

type ProfileRequest struct {
	UserId                  int                   `json:"userId" form:"userId"`
	Status                  string                `json:"status" form:"status"`
	IsSpecialist            *bool                 `json:"isSpecialist" form:"isSpecialist"`
	IdCard                  string                `json:"idCard" form:"idCard"`
	GovCard                 string                `json:"govCard" form:"govCard"`
	CardExpire              *time.Time            `json:"cardExpire" form:"cardExpire"`
	PrefixName              string                `json:"prefixName" form:"prefixName"`
	FirstName               string                `json:"firstName" form:"firstName"`
	LastName                string                `json:"lastName" form:"lastName"`
	BirthDate               *time.Time            `json:"birthDate" form:"birthDate"`
	Race                    string                `json:"race" form:"race"`
	Nation                  string                `json:"nation" form:"nation"`
	DomicileNo              string                `json:"domicileNo" form:"domicileNo"`
	DomicileMoo             string                `json:"domicileMoo" form:"domicileMoo"`
	DomicileSoi             string                `json:"domicileSoi" form:"domicileSoi"`
	DomicileRoad            string                `json:"domicileRoad" form:"domicileRoad"`
	DomicileProvince        int                   `json:"domicileProvince" form:"domicileProvince"`
	DomicileDistrict        int                   `json:"domicileDistrict" form:"domicileDistrict"`
	DomicileSubDistrict     int                   `json:"domicileSubDistrict" form:"domicileSubDistrict"`
	DomicileZipcode         string                `json:"domicileZipcode" form:"domicileZipcode"`
	DomicileTel             string                `json:"domicileTel" form:"domicileTel"`
	DomicileFax             string                `json:"domicileFax" form:"domicileFax"`
	DomicileEmail           string                `json:"domicileEmail" form:"domicileEmail"`
	AddressNo               string                `json:"addressNo" form:"addressNo"`
	AddressMoo              string                `json:"addressMoo" form:"addressMoo"`
	AddressSoi              string                `json:"addressSoi" form:"addressSoi"`
	AddressRoad             string                `json:"addressRoad" form:"addressRoad"`
	AddressSubDistrict      int                   `json:"addressSubDistrict" form:"addressSubDistrict"`
	AddressDistrict         int                   `json:"addressDistrict" form:"addressDistrict"`
	AddressProvince         int                   `json:"addressProvince" form:"addressProvince"`
	AddressZipcode          string                `json:"addressZipcode" form:"addressZipcode"`
	AddressTel              string                `json:"addressTel" form:"addressTel"`
	AddressFax              string                `json:"addressFax" form:"addressFax"`
	AddressEmail            string                `json:"addressEmail" form:"addressEmail"`
	ContactNo               string                `json:"contactNo" form:"contactNo"`
	ContactMoo              string                `json:"contactMoo" form:"contactMoo"`
	ContactSoi              string                `json:"contactSoi" form:"contactSoi"`
	ContactRoad             string                `json:"contactRoad" form:"contactRoad"`
	ContactSubDistrict      int                   `json:"contactSubDistrict" form:"contactSubDistrict"`
	ContactDistrict         int                   `json:"contactDistrict" form:"contactDistrict"`
	ContactProvince         int                   `json:"contactProvince" form:"contactProvince"`
	ContactZipcode          string                `json:"contactZipcode" form:"contactZipcode"`
	ContactTel              string                `json:"contactTel" form:"contactTel"`
	ContactFax              string                `json:"contactFax" form:"contactFax"`
	ContactEmail            string                `json:"contactEmail" form:"contactEmail"`
	WorkOccupation          string                `json:"workOccupation" form:"workOccupation"`
	WorkPosition            string                `json:"workPosition" form:"workPosition"`
	WorkPlaces              string                `json:"workPlaces" form:"workPlaces"`
	WorkRoad                string                `json:"workRoad" form:"workRoad"`
	WorkSubDistrict         int                   `json:"workSubDistrict" form:"workSubDistrict"`
	WorkDistrict            int                   `json:"workDistrict" form:"workDistrict"`
	WorkProvince            int                   `json:"workProvince" form:"workProvince"`
	WorkZipcode             string                `json:"workZipcode" form:"workZipcode"`
	WorkTel                 string                `json:"workTel" form:"workTel"`
	WorkFax                 string                `json:"workFax" form:"workFax"`
	WorkEmail               string                `json:"workEmail" form:"workEmail"`
	BossFirstName           string                `json:"bossFirstName" form:"bossFirstName"`
	BossLastName            string                `json:"bossLastName" form:"bossLastName"`
	BossNo                  string                `json:"bossNo" form:"bossNo"`
	BossMoo                 string                `json:"bossMoo" form:"bossMoo"`
	BossSoi                 string                `json:"bossSoi" form:"bossSoi"`
	BossRoad                string                `json:"bossRoad" form:"bossRoad"`
	BossSubDistrict         int                   `json:"bossSubDistrict" form:"bossSubDistrict"`
	BossDistrict            int                   `json:"bossDistrict" form:"bossDistrict"`
	BossProvince            int                   `json:"bossProvince" form:"bossProvince"`
	BossZipcode             string                `json:"bossZipcode" form:"bossZipcode"`
	BossTel                 string                `json:"bossTel" form:"bossTel"`
	BossFax                 string                `json:"bossFax" form:"bossFax"`
	BossEmail               string                `json:"bossEmail" form:"bossEmail"`
	WorkExperience          string                `json:"workExperience" form:"workExperience"`
	SplTypeID               int                   `json:"splTypeID" form:"splTypeID"`
	SplSubTypeID            int                   `json:"splSubTypeID" form:"splSubTypeID"`
	RegisQualification      string                `json:"regisQualification" form:"regisQualification"`
	RegisDocument           string                `json:"regisDocument" form:"regisDocument"`
	RegisEver               string                `json:"regisEver" form:"regisEver"`
	RegisEverYear           string                `json:"regisEverYear" form:"regisEverYear"`
	RegisEverPass           string                `json:"regisEverPass" form:"regisEverPass"`
	RegisEverPassNo         string                `json:"regisEverPassNo" form:"regisEverPassNo"`
	RegisEverNopass         string                `json:"regisEverNopass" form:"regisEverNopass"`
	RegisEverNopassDesc     string                `json:"regisEverNopassDesc" form:"regisEverNopassDesc"`
	FileAttachIdcard        *multipart.FileHeader `json:"fileAttachIdcard" form:"fileAttachIdcard"`
	FileAttachHouse         *multipart.FileHeader `json:"fileAttachHouse" form:"fileAttachHouse"`
	FileAttachGovCard       *multipart.FileHeader `json:"fileAttachGovCard" form:"fileAttachGovCard"`
	FileAttachQualification *multipart.FileHeader `json:"fileAttachQualification" form:"fileAttachQualification"`
	Remark                  string                `json:"remark" form:"remark"`
}

func (p Profile) Serialize() map[string]interface{} {

	return map[string]interface{}{
		"prefixName":    p.PrefixName,
		"firstName":     p.FirstName,
		"lastName":      p.LastName,
		"isSpecialist":  p.IsSpecialist,
		"idCard":        p.IdCard,
		"statusReqform": p.StatusReqform,
	}
}

func (p Profile) SerializeList() map[string]interface{} {
	return map[string]interface{}{
		"firstName":  p.FirstName,
		"lastName":   p.LastName,
		"email":      p.ContactEmail,
		"SplType":    p.SplSubType.SplType.NameTH,
		"SplSubType": p.SplSubType.NameTH,
	}
}
