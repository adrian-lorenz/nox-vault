package app

import (
	"errors"
	"github.com/adrian-lorenz/nox-vault/Middleware"
	"github.com/adrian-lorenz/nox-vault/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func CreateApp(c *gin.Context) {
	tokenInfo, errT := c.MustGet("TokenInfo").(Middleware.TokenData)
	if !errT {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	type sPUT struct {
		Name string `json:"name"  binding:"required"`
	}
	var s sPUT
	errB := c.ShouldBindJSON(&s)
	if errB != nil {
		c.JSON(500, gin.H{"message": "bad payload"})
		return
	}

	//create app
	a := database.App{
		Name:        s.Name,
		UUID:        uuid.New().String(),
		CreatorUUID: tokenInfo.UUID,
	}
	errCs := createApp(a)
	if errCs != nil {
		c.JSON(500, gin.H{"message": errCs.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "App created"})

}
func createApp(app database.App) error {
	var secExist database.App
	result := database.DB.Where(database.App{Name: app.Name}).First(&secExist)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = database.DB.Create(&app)
		if result.Error != nil {
			return errors.New("Fehler beim Erstellen der App: " + result.Error.Error())
		} else {
			return nil
		}
	} else if result.Error != nil {
		return errors.New("Fehler beim Suchen der App: " + result.Error.Error())

	} else {
		return errors.New("app existiert bereits")
	}
}
