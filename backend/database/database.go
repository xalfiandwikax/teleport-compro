package database

import (
	"backend/server/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB membuat koneksi ke database PostgreSQL
func ConnectDB(host string, port int, user, password, dbname string) (*gorm.DB, error) {
	// Format DSN (Data Source Name) untuk PostgreSQL
	// Format: host=... port=... user=... password=... dbname=... sslmode=...
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Buka koneksi ke database
	// gorm.Open membuka koneksi dengan driver PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err // Return error jika gagal connect
	}

	// Auto migrate: Buat/update tabel berdasarkan struct models
	// AutoMigrate akan:
	// 1. Buat tabel jika belum ada
	// 2. Tambah kolom jika ada field baru
	// 3. Tidak menghapus kolom yang sudah ada
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	// Buat superadmin default jika belum ada
	createDefaultSuperadmin(db)

	log.Println("Database connected successfully!")
	return db, nil
}

// createDefaultSuperadmin membuat user superadmin pertama
func createDefaultSuperadmin(db *gorm.DB) {
	var count int64

	// Hitung berapa banyak superadmin di database
	// SELECT COUNT(*) FROM users WHERE role = 'superadmin'
	db.Model(&models.User{}).Where("role = ?", models.RoleSuperadmin).Count(&count)

	// Jika belum ada superadmin sama sekali
	if count == 0 {
		superadmin := models.User{
			Name:     "Super Admin",
			Email:    "superadmin@mail.com",
			Password: "superadmin123", // INI HARUS DIUBAH DI PRODUCTION!
			Role:     models.RoleSuperadmin,
		}

		// Create user (BeforeCreate hook akan auto-hash password)
		if err := db.Create(&superadmin).Error; err == nil {
			log.Println("Default superadmin created:")
			log.Println("Email: superadmin@mail.com")
			log.Println("Password: superadmin123")
		}
	}
}
