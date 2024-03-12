package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"runtime"

	"github.com/adrian-lorenz/nox-vault/database"
	"github.com/adrian-lorenz/nox-vault/routes"
	log "github.com/sirupsen/logrus"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var GlobalDebug bool = false

func setupRouter() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	if os.Getenv("VAULT_DEBUG") == "true" {
		GlobalDebug = true
	}
	if _, exists := os.LookupEnv("VAULT_SECRET"); !exists {
		log.Panic("VAULT_SECRET not set")
	}

	log.Infof(
		"compiled for %s/%s\n",
		runtime.GOOS,
		runtime.GOARCH,
	)

	if err := database.ConnectDB(); err != nil {
		log.Panic("database error")
	}
	if !GlobalDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(cors.Default())
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.GET("/check", routes.CheckService)
	//router.GET("/senders/:ident", tools.CheckAcceptHeader, tools.TokenRequired(), routes.SenderGet)
	return router
	//hallo sandro
}

func main() {
	router := setupRouter()
	log.Infoln("Server started!")
	err := router.Run("0.0.0.0:" + os.Getenv("VAULT_PORT"))

	if err != nil {
		log.Error(err.Error())

		return
	}
}
