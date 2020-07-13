package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/judascrow/gomiddlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/judascrow/cojspcl-api/api/infrastructure"
	"github.com/judascrow/cojspcl-api/api/models"
	"github.com/judascrow/cojspcl-api/api/routes"
	"github.com/judascrow/cojspcl-api/api/seeds"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
	}
}

func drop(db *gorm.DB) {
	db.DropTableIfExists(
		&models.CasbinRule{},
		&models.Role{},
		&models.User{},
		&models.Province{},
		&models.District{},
		&models.SubDistrict{},
		&models.SplType{},
		&models.SplSubType{},
		&models.Profile{},
	)
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.CasbinRule{})
	db.AutoMigrate(&models.Province{})
	db.AutoMigrate(&models.District{})
	db.AutoMigrate(&models.SubDistrict{})
	db.AutoMigrate(&models.SplType{})
	db.AutoMigrate(&models.SplSubType{})
	db.AutoMigrate(&models.Profile{})
}

func addDbConstraints(db *gorm.DB) {

}

func create(db *gorm.DB) {
	drop(db)
	migrate(db)
	addDbConstraints(db)
}

func Run() {

	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	database := infrastructure.InitDb()
	defer database.Close()

	args := os.Args
	fmt.Println(args)
	if len(args) > 1 && args[1] != "main.go" {
		first := args[1]
		second := ""
		if len(args) > 2 {
			second = args[2]
		}

		if first == "create" {
			create(database)
		} else if first == "seed" {
			seeds.Seed()
			os.Exit(0)
		} else if first == "migrate" {
			migrate(database)
		}

		if second == "seed" {
			seeds.Seed()
			os.Exit(0)
		} else if first == "migrate" {
			migrate(database)
		}

		if first != "" && second == "" {
			os.Exit(0)
		}
	}

	migrate(database)

	gin.SetMode(os.Getenv("SERVER_RUN_MODE"))

	r := routes.InitRouter()
	defer gomiddlewares.CloseLogFile()

	port := os.Getenv("SERVER_PORT")
	readTimeoutInt, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	writeTimeoutInt, _ := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	readTimeout := time.Duration(readTimeoutInt) * time.Second
	writeTimeout := time.Duration(writeTimeoutInt) * time.Second
	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Print("API is ready to listen and serve on PORT : " + port)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

}
