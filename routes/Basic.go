package routes

import (
	"github.com/adrian-lorenz/nox-vault/globals"

	"github.com/adrian-lorenz/nox-vault/database"
	"github.com/gin-gonic/gin"
)

type Check struct {
	Database bool `json:"database"`
	MKey     bool `json:"mkey"`
	Looked   bool `json:"looked"`
}

func CheckService(c *gin.Context) {
	cfov := Check{}
	cfov.Looked = globals.Look
	if globals.MasterKey != "" {
		cfov.MKey = true
	} else {
		cfov.MKey = false
	}
	var count int64
	err := database.DB.Raw("SELECT 1").Count(&count).Error
	if err != nil {
		cfov.Database = false
	} else {
		cfov.Database = count == 1
	}
	c.JSON(200, cfov)

}
