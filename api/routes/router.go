package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/judascrow/cojspcl-api/api/controllers"
	"github.com/judascrow/gomiddlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	// Prometheus
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	p.ReqCntURLLabelMappingFn = MappingFn

	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"success": false, "error": "Method Not Allowed"})
		return
	})

	// middlewares
	r.Use(gomiddlewares.GoLogger(), gomiddlewares.GoCors())
	if os.Getenv("APP_ENV") == "dev" {
		gin.ForceConsoleColor()
		r.Use(gin.Logger())
	}
	r.Use(gin.Recovery())

	r.MaxMultipartMemory = 8 << 20
	// File Server
	r.Use(static.Serve("/upload", static.LocalFile("./upload", false)))

	// Routes
	apiv1 := r.Group(os.Getenv("APP_API_BASE_URL"))

	// Swagger
	apiv1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Healthcheck
	apiv1.GET("/healthcheck", Healthcheck)

	// Auth Middleware
	authMiddleware := AuthMiddlewareJWT()

	// Auth API
	auth := apiv1.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler) // Refresh time can be longer than token timeout
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.POST("/register", controllers.Register)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/me", controllers.GetUserMe)
	}

	// User API
	users := apiv1.Group("/users")
	users.Use(authMiddleware.MiddlewareFunc())
	{
		users.GET("", controllers.GetAllUsers)
		users.GET("/:slug", controllers.GetUserBySlug)
		users.POST("", controllers.CreateUser)
		users.PUT("/:slug", controllers.UpdateUser)
		users.DELETE("/:slug", controllers.DeleteUser)
		users.POST("/:slug/password", controllers.ChangePassword)
		users.POST("/:slug/avatar", controllers.UploadAvatar)
	}

	// Role API
	roles := apiv1.Group("/roles")
	roles.Use(authMiddleware.MiddlewareFunc())
	{
		roles.GET("", controllers.GetAllRoles)
	}

	// Address API

	apiv1.GET("/provinces", controllers.GetAllProvinces)
	apiv1.GET("/provinces/:id", controllers.GetProvinceByID)
	apiv1.GET("/districts/:id", controllers.GetDistrictByID)
	apiv1.GET("/subdistricts/:id", controllers.GetSubDistrictByID)

	provinces := apiv1.Group("/province")
	provinces.GET("/:provinceID/districts", controllers.GetDistrictsByProvinceID)
	provinces.GET("/:provinceID/district/:districtID/subDistricts", controllers.GetSubDistrictsByDistrictID)

	// Speciallist Type API
	splTypes := apiv1.Group("/spltypes")
	splTypes.GET("", controllers.GetAllSplTypes)
	apiv1.GET("/spltype/:id", controllers.GetSplTypeByID)
	splTypes.GET("/:splTypeID/splsubtypes", controllers.GetSplSubTypesBySplTypeID)
	apiv1.GET("/splsubtype/:id", controllers.GetSplSubTypeByID)

	// Reqform API
	reqform := apiv1.Group("/reqforms")
	reqform.Use(authMiddleware.MiddlewareFunc())
	{
		reqform.POST("", controllers.CreateProfile)
	}

	return r
}

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "API is Online",
	})

}

func MappingFn(c *gin.Context) string {
	url := c.Request.URL.Path
	for _, p := range c.Params {
		if p.Key == "id" {
			url = strings.Replace(url, p.Value, ":id", 1)
			break
		}
	}
	return url
}
