package config

import (
	user "clean_arch/features/user/entity"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	SERVER_PORT int
	DB_USER 	string
	DB_PASS 	string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	SECRET      string
}

func InitDB() *gorm.DB {

	config := loadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	return db
}

func loadConfig() Config {
	godotenv.Load(".env")
	
	serverPort, err := strconv.Atoi(os.Getenv("SERVER"))

	if err != nil {
		panic(err)
	}
	
	return Config{
		SERVER_PORT: serverPort,
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_NAME: os.Getenv("DB_NAME"),
		SECRET:  os.Getenv("SECRET"),
	}
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&user.UserEntity{})
}