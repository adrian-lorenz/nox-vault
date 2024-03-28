package secrets

import (
	"errors"
	"github.com/adrian-lorenz/nox-vault/Middleware"
	"github.com/adrian-lorenz/nox-vault/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func AddSecret(c *gin.Context) {
	tokenInfo, errT := c.MustGet("TokenInfo").(Middleware.TokenData)
	if !errT {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	type sPUT struct {
		Name   string `json:"name"  binding:"required"`
		Secret string `json:"content" binding:"required"`
		App    string `json:"app"  binding:"required"`
	}
	var s sPUT
	errB := c.ShouldBindJSON(&s)
	if errB != nil {
		c.JSON(500, gin.H{"message": "bad payload"})
		return
	}
	cGo := checkUserApp(tokenInfo.UUID, s.App)

	if !cGo {
		c.JSON(500, gin.H{})
		return
	}
	//create secret
	secret := database.Secret{
		Name:         s.Name,
		UUID:         uuid.New().String(),
		Content:      "This is a test",
		AppUUID:      s.App,
		CreatorUUID:  tokenInfo.UUID,
		ModifierUUID: tokenInfo.UUID,
	}
	errCs := createSecret(secret)
	if errCs != nil {
		c.JSON(500, gin.H{"message": errCs.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Test secret created"})
}

func createSecret(secret database.Secret) error {
	var secExist database.Secret
	result := database.DB.Where(database.Secret{Name: secret.Name, AppUUID: secret.AppUUID}).First(&secExist)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = database.DB.Create(&secret)
		if result.Error != nil {
			return errors.New("Fehler beim Erstellen des Secrets: " + result.Error.Error())
		} else {
			return nil
		}
	} else if result.Error != nil {
		return errors.New("Fehler beim Suchen des Secrets: " + result.Error.Error())

	} else {
		return errors.New("secret existiert bereits")
	}
}

func checkUserApp(userUUID string, appUUID string) bool {
	var userApp database.UserFusion
	result := database.DB.Where(database.UserFusion{UserUUID: userUUID, AppUUID: appUUID}).First(&userApp)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	} else if result.Error != nil {
		return false
	} else {
		return true
	}
}
