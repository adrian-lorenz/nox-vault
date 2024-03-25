package routes

import (
	"fmt"

	"github.com/adrian-lorenz/nox-vault/cmux"
	"github.com/adrian-lorenz/nox-vault/database"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Check struct {
	Database bool   `json:"database"`
	Status   bool   `json:"status"`
	Message  string `json:"message"`
}

func CheckService(c *gin.Context) {
	log.Info("GET route check")
	cfov := Check{}
	var count int64
	err := database.DB.Raw("SELECT 1").Count(&count).Error
	if err != nil {
		fmt.Println(err)
		cfov.Database = false
		cfov.Status = false
		cfov.Message = "DB Connect Error"
	} else {
		cfov.Database = count == 1
		cfov.Status = true
		cfov.Message = "ok"
	}
	c.JSON(200, cmux.JSON{"message": "Hello World"})

}
