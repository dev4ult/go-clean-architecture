package entity

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	ID int `gorm:"type:int(11)"`
	Username string `gorm:"varchar(255)"`
	Email string `gorm:"varchar(255)"`
	Password string `gorm:"varchar(64)"`
}