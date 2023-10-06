package config

import (
	user "clean_arch/features/user/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := "root@tcp(127.0.0.1:3306)/try_middleware?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.UserEntity{})
}