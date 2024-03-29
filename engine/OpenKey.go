package engine

import (
	"github.com/adrian-lorenz/nox-vault/cfernet"
	"github.com/adrian-lorenz/nox-vault/database"
	"github.com/adrian-lorenz/nox-vault/globals"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//hDHrEfAXkF_CqUCWHDRVYPH71nekdBiny28ELjGGECc=

func OpenKey(c *gin.Context) {
	type sKey struct {
		Key string `json:"key"  binding:"required"`
	}
	var s sKey
	errB := c.ShouldBindJSON(&s)
	if errB != nil {
		c.JSON(500, gin.H{"message": "bad payload"})
		return
	}
	if globals.JWTKey != "" || globals.MasterKey != "" || !globals.Look {
		c.JSON(401, gin.H{})
		return
	}
	var settings database.Settings
	result := database.DB.First(&settings)
	if result.Error != nil {
		log.Error("Fehler beim Lesen der Einstellungen:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fehler beim Lesen der Einstellungen"})
		return
	}
	if settings.TestKey == "" {
		log.Error("Test-Key nicht vorhanden")
		c.JSON(http.StatusConflict, gin.H{"message": "Test-Key nicht vorhanden"})
		return
	}
	cfern := cfernet.NewEncryptor(s.Key)
	decKey := cfern.Decrypt(settings.TestKey)
	if decKey != "nox-vault" {
		log.Error("Key mismatch")
		c.JSON(http.StatusConflict, gin.H{"message": "Key mismatch"})
		return
	} else {
		globals.MasterKey = s.Key
		globals.JWTKey = s.Key
		globals.Look = false
		c.JSON(http.StatusOK, gin.H{"message": "unlooked"})
	}

}
