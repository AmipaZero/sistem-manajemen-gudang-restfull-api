package config

import (
	"log"
	"sistem-manajemen-gudang/model/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:mipa@tcp(127.0.0.1:3306)/sbp?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi DB: ", err)
	}

	// Auto migrate
	DB.AutoMigrate(&domain.Product{})
	DB.AutoMigrate(&domain.Inbound{})
	DB.AutoMigrate(&domain.Outbound{})
	DB.AutoMigrate(&domain.User{})
}