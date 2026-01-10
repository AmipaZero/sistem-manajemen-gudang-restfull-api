package main

import (
	"log"

	"sistem-manajemen-gudang/config"
	"sistem-manajemen-gudang/app"
	"sistem-manajemen-gudang/util"
)

func main() {
	// Load environment variable
	util.InitEnv()

	// Connect database
	config.ConnectDB()
	db := config.DB
	if db == nil {
		log.Fatal("❌ Gagal koneksi ke database")
	}
	log.Println("✅ Database connected")

	// Setup router
	r := app.SetupRouter(db)

	// Run server
	if err := r.Run(":3000"); err != nil {
		log.Fatal("❌ Gagal menjalankan server:", err)
	}
}