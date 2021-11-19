package config

import (
	"fmt"
	"mvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const JWT_SECRET string = "testmvc"

const JWT_EXP int = 1

func InitDB(status string) {
	db := "gormdbtry"
	if status == "testing" {
		db = "gormdbtry-test"
	}
	connectionString := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?parseTime=True", db)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	MigrateDB()
}

func MigrateDB() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Admin{})
	DB.AutoMigrate(&models.CreditCard{})
	DB.AutoMigrate(&models.Kabs{})
}
