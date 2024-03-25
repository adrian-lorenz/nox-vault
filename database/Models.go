package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null"`
	Password string `gorm:"size:255;not null"`
	UUID     string `gorm:"size:255;primaryKey;unique"`
	Email    string `gorm:"size:255;not null"`
	Role     string `gorm:"size:100;not null"`
}
type App struct {
	gorm.Model
	Name string `gorm:"size:100;not null"`
	UUID string `gorm:"size:255;primaryKey"` // UUID als Primärschlüssel

}
type UserFusion struct {
	gorm.Model
	AppUUID  string `gorm:"type:uuid;not null"`
	UserUUID string `gorm:"type:uuid;not null"`
}

type Secret struct {
	gorm.Model
	Name         string `gorm:"size:100;not null"`
	Content      string `gorm:"size:255;not null"`
	UUID         string `gorm:"size:255;primaryKey"`
	AppUUID      string `gorm:"type:uuid;not null"`
	CreatorUUID  string `gorm:"type:uuid;not null"`
	ModifierUUID string `gorm:"type:uuid;not null"`
}
