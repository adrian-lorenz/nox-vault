package engine

import (
	"errors"
	"github.com/adrian-lorenz/nox-vault/cfernet"
	"github.com/adrian-lorenz/nox-vault/database"
	"github.com/adrian-lorenz/nox-vault/security"
	"github.com/adrian-lorenz/nox-vault/tools"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func CreateKey(c *gin.Context) {

	iip := tools.GetIP(c)
	if !security.CheckWhitelistsInt(iip) {
		c.JSON(http.StatusUnauthorized, gin.H{})
		log.Error("IP not whitelisted", iip)
		c.Abort()
		return
	}

	var settings database.Settings

	result := database.DB.First(&settings)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Error("Fehler beim Lesen der Einstellungen:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fehler beim Lesen der Einstellungen"})
		return
	}
	if settings.TestKey != "" {
		log.Error("Test-Key bereits vorhanden")
		c.JSON(http.StatusConflict, gin.H{"message": "Test-Key bereits vorhanden"})
		return
	}
	//gen key
	key, err := cfernet.CreateFernetKey(32)
	if err != nil {
		log.Error("Fehler beim Generieren des Schlüssels:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fehler beim Generieren des Schlüssels"})
		return
	}

	frnt := cfernet.NewEncryptor(key)
	encKey := frnt.Encrypt("nox-vault")
	nsett := database.Settings{TestKey: encKey}
	result = database.DB.Save(&nsett)
	if result.Error != nil {
		log.Error("Fehler beim Speichern des Schlüssels:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fehler beim Speichern des Schlüssels"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Test-Key erstellt", "key": key})

}
