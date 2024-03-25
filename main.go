package main

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

	err := godotenv.Load()
	if err != nil {
		log.Infoln("No environment file found - set PROD")
		gin.SetMode(gin.ReleaseMode)
		globals.Mode = "PROD"
	} else {
		globals.Mode = os.Getenv("MODE")
		log.Infoln("Environment file found - set " + globals.Mode)
	}

	if os.Getenv("VAULT_DSN") == "" {
		log.Fatal("DATABASE environment variable not set")
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
	// routes
	router.GET("/check", routes.CheckService)

	router.GET("/test", func(context *gin.Context) {
		//search for user
		var user database.User
		resA := database.DB.Where(database.User{Username: "admin"}).First(&user)
		if resA.Error != nil {
			log.Error(resA.Error)
		}
		var app database.App
		resB := database.DB.Where(database.App{Name: "TestApp"}).First(&app)
		if resB.Error != nil {
			log.Error(resB.Error)
			return // Fügt eine Rückgabe hinzu, um bei einem Fehler zu stoppen
		}

		//create secret
		secret := database.Secret{
			Name:         "TestSecret",
			UUID:         uuid.New().String(),
			Content:      "This is a test",
			AppUUID:      app.UUID,
			CreatorUUID:  user.UUID,
			ModifierUUID: user.UUID,
		}
		var secExist database.Secret
		result := database.DB.Where(database.Secret{Name: secret.Name, AppUUID: secret.AppUUID}).First(&secExist)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Datensatz nicht gefunden, also erstellen wir einen neuen
			result = database.DB.Create(&secret)
			if result.Error != nil {
				log.Printf("Fehler beim Erstellen des Secrets: %v", result.Error)
			} else {
				log.Println("Secret erfolgreich erstellt.")
			}
		} else if result.Error != nil {
			// Ein anderer Fehler ist aufgetreten
			log.Printf("Fehler beim Suchen des Secrets: %v", result.Error)
		} else {
			// Secret existiert bereits, keine Aktion erforderlich
			log.Println("Secret existiert bereits, keine Erstellung notwendig.")
		}
		context.JSON(200, gin.H{"message": "Test secret created"})
	})

	log.Infoln("Starting server on localhost:5050")

	ErrM := router.Run(":" + os.Getenv("VAULT_PORT"))
	if ErrM != nil {
		log.Fatal(ErrM)
	}
}
