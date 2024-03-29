package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null"`
	Password string `gorm:"size:255;not null"`
	UUID     string `gorm:"size:255;unique"`
	Email    string `gorm:"size:255;not null"`
	Role     string `gorm:"size:100;not null"`
}
type App struct {
	gorm.Model
	Name string `gorm:"size:100;not null"`
	UUID string `gorm:"size:255;unique"`
}
type UserFusion struct {
	gorm.Model
	AppUUID  string `gorm:"size:255;not null"`
	UserUUID string `gorm:"size:255;not null"`
}

type Secret struct {
	gorm.Model
	Name         string `gorm:"size:100;not null"`
	Content      string `gorm:"size:255;not null"`
	UUID         string `gorm:"size:255;unique"`
	AppUUID      string `gorm:"size:255;not null"`
	CreatorUUID  string `gorm:"size:255;not null"`
	ModifierUUID string `gorm:"size:255;not null"`
	App          App    `gorm:"foreignKey:AppUUID;references:UUID"`
	Owner        User   `gorm:"foreignKey:CreatorUUID;references:UUID"`
}

type Settings struct {
	gorm.Model
	TestKey string `gorm:"size:255;not null"`
	Guid    string `gorm:"size:255;not null"`
}
