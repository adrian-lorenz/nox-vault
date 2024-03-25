package database

import (
	"errors"
	"os"

	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB //

func ConnectDB() error {
	db, err := gorm.Open(sqlite.Open(os.Getenv("VAULT_DSN")+"?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return errors.New("can't connect to database")
	}
	errCre := db.AutoMigrate(&User{}, &App{}, &UserFusion{}, &Secret{})
	if errCre != nil {
		return errCre
	}

	DB = db
	prepareData(DB)
	return nil
}

func prepareData(DB *gorm.DB) {
	// Passwort verschlüsseln
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("muha"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Fehler beim Generieren des verschlüsselten Passworts: %v", err)
		return
	}

	// Benutzer erstellen oder finden, falls er bereits existiert
	user := User{
		Username: "admin",
		Password: string(hashedPassword),
		Email:    "test@test.de",
		Role:     "admin",
		UUID:     "7A4C2E09-207C-466E-8814-C902FB296432",
	}
	result := DB.FirstOrCreate(&user)
	if result.Error != nil {
		log.Printf("Fehler beim Erstellen oder Finden des Benutzers: %v", result.Error)
		return
	}

	// App erstellen oder finden, falls sie bereits existiert

	app := App{
		Name: "TestApp",
		UUID: "A3DAD5CB-942B-4939-A13D-979C3C4F7384",
	}
	result = DB.FirstOrCreate(&app)
	if result.Error != nil {
		log.Printf("Fehler beim Erstellen oder Finden der App: %v", result.Error)
		return
	}

	// Add Fusion
	userFusion := UserFusion{
		AppUUID:  app.UUID,
		UserUUID: user.UUID,
	}
	result = DB.Where(UserFusion{AppUUID: app.UUID, UserUUID: user.UUID}).FirstOrCreate(&userFusion)
	if result.Error != nil {
		log.Printf("Fehler beim Erstellen oder Finden der Fusion: %v", result.Error)
		return
	}

	return

}
