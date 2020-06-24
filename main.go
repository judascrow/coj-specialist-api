package main

import (
	"github.com/judascrow/cojspcl-api/api"
	_ "github.com/judascrow/cojspcl-api/docs"
)

// @title COJ Specialist API
// @version 1.0
// @description Rest API document <style>.models {display: none !important}</style>
// @termsOfService COJ

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	api.Run()

}
