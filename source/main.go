package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"termdo.com/gateway-api/source/app/config"
	"termdo.com/gateway-api/source/core/auth"
	"termdo.com/gateway-api/source/core/tasks"
)

func main() {
	// Flags
	devMode := flag.Bool("dev", false, "Enable development mode")
	flag.Parse()

	// Load environment variables
	if *devMode {
		gin.SetMode(gin.DebugMode)
		godotenv.Load(".env")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Load configuration
	config.LoadAppConfig()
	config.LoadApiConfig()

	// Application
	app := gin.Default()
	app.SetTrustedProxies(nil)

	// Routes
	auth.BuildRoutes(app)
	tasks.BuildRoutes(app)

	// Start
	log.Fatal(app.Run(":" + strconv.Itoa(config.AppPort)))
}
