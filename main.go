package main

import (
	"github.com/adrian-lorenz/nox-vault/Middleware"
	"github.com/adrian-lorenz/nox-vault/app"
	"github.com/adrian-lorenz/nox-vault/engine"
	"github.com/adrian-lorenz/nox-vault/secrets"
	"github.com/adrian-lorenz/nox-vault/security"
	"github.com/adrian-lorenz/nox-vault/vault"
	"os"

	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/adrian-lorenz/nox-vault/database"
	"github.com/adrian-lorenz/nox-vault/globals"
	"github.com/adrian-lorenz/nox-vault/routes"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	dirs := []string{"noxvault/storage", "noxvault/certs", "noxvault/config"}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			errMk := os.MkdirAll(dir, 0755)
			if errMk != nil {
				return
			}
		}
	}
	//check mkdir storage

	err := godotenv.Load()
	if err != nil {
		log.Infoln("No environment file found - set PROD")
		gin.SetMode(gin.ReleaseMode)
		globals.Mode = "PROD"
	} else {
		globals.Mode = os.Getenv("MODE")
		log.Infoln("Environment file found - set " + globals.Mode)
	}

	if os.Getenv("VAULT_PORT") == "" {
		log.Fatal("environment variables not set")
	}

	log.Infof(
		"compiled for %s/%s\n",
		runtime.GOOS,
		runtime.GOARCH,
	)

	if err := database.ConnectDB(); err != nil {
		log.Panic("database error")
	} else {
		log.Info("Database connected")
	}

	//cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"*"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	// STUFF
	router.GET("/check", routes.CheckService)
	//INTERNAL
	router.POST("/secret/create", Middleware.TokenRequiredLst(globals.Internal), secrets.AddSecret)
	router.POST("/secret/update", Middleware.TokenRequiredLst(globals.Internal))
	router.POST("/app/create", Middleware.TokenRequiredLst(globals.Internal), app.CreateApp)
	router.GET("/key/gen", Middleware.SysWhitelist(), engine.CreateKey)
	router.POST("/vault/open", Middleware.SysWhitelist(), vault.OpenVault)
	router.POST("/login", Middleware.SysWhitelist(), security.Login)

	//EXTERNAL
	router.POST("/secret/get", Middleware.TokenRequiredLst(globals.Read))

	log.Infoln("Starting server on localhost:5050")

	ErrM := router.Run(":" + os.Getenv("VAULT_PORT"))
	if ErrM != nil {
		log.Fatal(ErrM)
	}
}
