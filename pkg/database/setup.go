package database

import (
	"fmt"
	"log"

	"chitchat/app/models"

	"chitchat/pkg/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.GetEnv("DB_USER", ""),
		env.GetEnv("DB_PASSWORD", ""),
		env.GetEnv("DB_HOST", "127.0.0.1"),
		env.GetEnv("DB_PORT", "3306"),
		env.GetEnv("DB_NAME", ""),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database \n", err.Error())
		panic(err)
	}
	DB.Logger = logger.Default.LogMode(logger.Info)
	err = DB.AutoMigrate(&models.User{}, &models.UserSession{})
	if err != nil {
		log.Fatal("failed to migrate database \n", err.Error())
		panic(err)
	}
	fmt.Println("Successfully migrated database")
}
