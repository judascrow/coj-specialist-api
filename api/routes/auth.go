package routes

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/judascrow/cojspcl-api/api/infrastructure"
	"github.com/judascrow/cojspcl-api/api/models"
	"github.com/judascrow/cojspcl-api/api/utils/messages"
	"github.com/judascrow/gomiddlewares/jwt"

	"golang.org/x/crypto/bcrypt"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var casbinEnforcer *casbin.Enforcer
var db *gorm.DB
var err error

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User struct alias
type User = models.User

var identityKey = "slug"
var identityUsername = "username"
var identityRole = "roleId"

// @Summary เข้าสู่ระบบ
// @Description เข้าสู่ระบบ
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param login body login true "ลงชื่อเข้าสู่ระบบ"
// @Success 200 {object} models.SwagLogin
// @Router /auth/login [post]
func AuthMiddlewareJWT() *jwt.GinJWTMiddleware {
	db := infrastructure.GetDB()

	Dbdriver := os.Getenv("DB_DRIVER")
	DbName := os.Getenv("DB_NAME")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")

	pg_conn_info := DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8&parseTime=True&loc=Local"
	casbin_adapter := gormadapter.NewAdapter(Dbdriver, pg_conn_info, true)
	e := casbin.NewEnforcer("./auth.conf", casbin_adapter)
	casbinEnforcer = e
	e.LoadPolicy()

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey:      v.Slug,
					identityUsername: v.Username,
					identityRole:     v.RoleID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			return &User{
				Slug:     claims["slug"].(string),
				Username: claims["username"].(string),
				RoleID:   int(claims["roleId"].(float64)),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.BindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			var user User
			if err := db.Set("gorm:auto_preload", true).Where("username = ? AND status = 'A' ", username).First(&user).Error; err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			if checkHash(password, user.Password) {
				return &User{
					Slug:     user.Slug,
					Username: user.Username,
					RoleID:   user.RoleID,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			//fmt.Println(len(data.(*User).UserRoles))
			if v, ok := data.(*User); ok && v.Username != "" && v.RoleID != 0 {
				v0 := strconv.Itoa(v.RoleID)

				return casbinEnforcer.Enforce(v0, c.Request.URL.String(), c.Request.Method)

			}

			return false
		},
		LoginResponse: func(c *gin.Context, code int, token string, t time.Time, claims map[string]interface{}) {
			var user models.User
			slug := claims[identityKey]
			if slug != "" {
				db := infrastructure.GetDB()
				db.Set("gorm:auto_preload", true).Where("slug = ?", slug).First(&user)
			}
			c.JSON(http.StatusOK, gin.H{
				"token":   token,
				"expire":  t.Format(time.RFC3339),
				"success": true,
				"message": messages.Logged,
				"data":    user.Serialize(),
			})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"success": false,
				"message": "ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง",
				"data":    map[string]interface{}{},
			})
		},

		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}

func checkHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
