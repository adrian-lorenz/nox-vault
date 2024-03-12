package database

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //

func ConnectDB() error {
	db, err := gorm.Open(mysql.Open(os.Getenv("VAULT_DSN")), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return errors.New("can't connect to database")
	}
	errCre := db.AutoMigrate(&Store{},&Keys{},&Fusion{})
	if errCre != nil {
		return errCre
	}

	DB = db
	return nil
}
