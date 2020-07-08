package seeds

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"
	"github.com/judascrow/cojspcl-api/api/infrastructure"
	"github.com/judascrow/cojspcl-api/api/models"
	"golang.org/x/crypto/bcrypt"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var StringNumberRunes = []rune("1234567890")

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randomStringNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = StringNumberRunes[rand.Intn(len(StringNumberRunes))]
	}
	return string(b)
}

func seedAdmin(db *gorm.DB) {
	count := 0
	adminRole := models.Role{Name: "admin", NameTH: "ผู้ดูแลระบบ", Description: "Only for admin"}
	query := db.Model(&models.Role{}).Where("name = ?", "admin")
	query.Count(&count)

	if count == 0 {
		db.Create(&adminRole)
	} else {
		query.First(&adminRole)
	}

	adminRoleUsers := 0
	var adminUsers []models.User
	db.Model(&adminRole).Related(&adminUsers, "Users")

	db.Model(&models.User{}).Where("username = ?", "admin").Count(&adminRoleUsers)
	if adminRoleUsers == 0 {

		// query.First(&adminRole) // First would fetch the Role admin because the query status name='ROLE_ADMIN'
		password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		// Approach 1
		user := models.User{FirstName: "AdminFN", LastName: "AdminLN", Email: "admin@golang.com", Username: "admin", Password: string(password), RoleID: 1}

		// Do not try to update the adminRole
		db.Set("gorm:association_autoupdate", false).Create(&user)

		if db.Error != nil {
			print(db.Error)
		}
	}
}

func seedStaff(db *gorm.DB) {
	count := 0
	staffRole := models.Role{Name: "staff", NameTH: "เจ้าหน้าที่", Description: "Only for staff"}
	query := db.Model(&models.Role{}).Where("name = ?", "staff")
	query.Count(&count)

	if count == 0 {
		db.Create(&staffRole)
	} else {
		query.First(&staffRole)
	}

	staffRoleUsers := 0
	var staffUsers []models.User
	db.Model(&staffRole).Related(&staffUsers, "Users")

	db.Model(&models.User{}).Where("username = ?", "staff").Count(&staffRoleUsers)
	if staffRoleUsers == 0 {

		// query.First(&adminRole) // First would fetch the Role admin because the query status name='ROLE_ADMIN'
		password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		// Approach 1
		user := models.User{FirstName: "StaffFN", LastName: "StaffLN", Email: "staff@golang.com", Username: "staff", Password: string(password), RoleID: 2}

		// Do not try to update the adminRole
		db.Set("gorm:association_autoupdate", false).Create(&user)

		if db.Error != nil {
			print(db.Error)
		}
	}
}

func seedUsers(db *gorm.DB) {
	count := 0
	role := models.Role{Name: "user", NameTH: "ผู้ใช้งานทั่วไป", Description: "Only for standard users"}
	q := db.Model(&models.Role{}).Where("name = ?", "user")
	q.Count(&count)

	if count == 0 {
		db.Create(&role)
	} else {
		q.First(&role)
	}

	var standardUsers []models.User
	db.Model(&role).Related(&standardUsers, "Users")
	usersCount := len(standardUsers)
	usersToSeed := 5
	usersToSeed -= usersCount
	if usersToSeed > 0 {
		for i := 0; i < usersToSeed; i++ {
			password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
			user := models.User{FirstName: fake.FirstName(), LastName: fake.LastName(), Email: fake.EmailAddress(), Username: fake.UserName(),
				Password: string(password), RoleID: 3}
			// No need to add the role as we did for seedAdmin, it is added by the BeforeSave hook
			db.Set("gorm:association_autoupdate", false).Create(&user)
		}
	}
}

func seedCasbinRule(db *gorm.DB) {
	var casbinRule [3]models.CasbinRule

	db.Where(&models.CasbinRule{PType: "p", V0: "1", V1: "/api/v1/*"}).Attrs(models.CasbinRule{V2: "(GET)|(POST)|(PUT)|(DELETE)"}).FirstOrCreate(&casbinRule[0])
	db.Where(&models.CasbinRule{PType: "p", V0: "2", V1: "/api/v1/*"}).Attrs(models.CasbinRule{V2: "(GET)|(POST)|(PUT)|(DELETE)"}).FirstOrCreate(&casbinRule[1])
	db.Where(&models.CasbinRule{PType: "p", V0: "3", V1: "/api/v1/auth/me"}).Attrs(models.CasbinRule{V2: "(GET)|(PUT)"}).FirstOrCreate(&casbinRule[2])

}

func seedProvince(db *gorm.DB) {
	file, err := ioutil.ReadFile("./data/provinces.json")
	if err != nil {
		panic(err)
	}
	provinces := []models.Province{}

	err = json.Unmarshal([]byte(file), &provinces)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(provinces); i++ {
		db.Where(&provinces[i]).FirstOrCreate(&models.Province{})
	}
}

func seedDistrict(db *gorm.DB) {
	file, err := ioutil.ReadFile("./data/districts.json")
	if err != nil {
		panic(err)
	}
	districts := []models.District{}

	err = json.Unmarshal([]byte(file), &districts)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(districts); i++ {
		db.Where(&districts[i]).FirstOrCreate(&models.District{})
	}
}

func seedSubDistricts(db *gorm.DB) {
	file, err := ioutil.ReadFile("./data/subdistricts.json")
	if err != nil {
		panic(err)
	}
	subDistricts := []models.SubDistrict{}

	err = json.Unmarshal([]byte(file), &subDistricts)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(subDistricts); i++ {
		db.Where(&subDistricts[i]).FirstOrCreate(&models.SubDistrict{})
	}
}

func seedSplTypes(db *gorm.DB) {
	file, err := ioutil.ReadFile("./data/spltypes.json")
	if err != nil {
		panic(err)
	}
	splTypes := []models.SplType{}

	err = json.Unmarshal([]byte(file), &splTypes)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(splTypes); i++ {
		db.Where(&splTypes[i]).FirstOrCreate(&models.SplType{})
	}
}

func seedSplSubTypes(db *gorm.DB) {
	file, err := ioutil.ReadFile("./data/splsubtypes.json")
	if err != nil {
		panic(err)
	}
	splSubTypes := []models.SplSubType{}

	err = json.Unmarshal([]byte(file), &splSubTypes)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(splSubTypes); i++ {
		db.Where(&splSubTypes[i]).FirstOrCreate(&models.SplSubType{})
	}
}

func Seed() {
	db := infrastructure.GetDB()
	rand.Seed(time.Now().UnixNano())
	seedAdmin(db)
	seedStaff(db)
	seedUsers(db)
	seedCasbinRule(db)
	seedProvince(db)
	seedDistrict(db)
	seedSubDistricts(db)
	seedSplTypes(db)
	seedSplSubTypes(db)
}
