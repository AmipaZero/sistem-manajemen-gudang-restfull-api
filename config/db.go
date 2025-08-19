package config

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sistem-manajemen-gudang/model"
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
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.Inbound{})
	DB.AutoMigrate(&model.Outbound{})
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Supplier{})
}
// Sistem Manajemen Gudang Perusahaan
// +----------------+          +-----------------+          +------------------+
// |   products     |◄────────▶|    inbounds     |          |   outbounds      |
// +----------------+          +-----------------+          +------------------+
// | id (PK)        |          | id (PK)         |          | id (PK)          |
// | name           |          | product_id (FK) |─────────▶| product_id (FK)  |
// | sku (unique)   |          | quantity        |          | quantity         |
// | category       |          | received_at     |          | sent_at          |
// | unit           |          | supplier        |          | destination      |
// | stock          |          +-----------------+          +------------------+
// +----------------+

// Relasi:
// - Satu product bisa punya banyak barang masuk (inbounds)
// - Satu product bisa punya banyak barang keluar (outbounds)
// id (PK)
// username
// password_hash
// role          -- enum: admin / staff
// created_at
// id (PK)
// name
// contact_person
// phone
// address