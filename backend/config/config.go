package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config struct menyimpan semua konfigurasi aplikasi
type Config struct {
	DBHost     string // localhost
	DBPort     int    // 5432
	DBUser     string // postgres
	DBPassword string // password123
	DBName     string // auth_rbac
	JWTSecret  string // secret key untuk JWT
	Port       string // 8080
}

// LoadConfig membaca konfigurasi dari file .env
func LoadConfig() (*Config, error) {
	// 1. Baca file .env
	// godotenv.Load() membaca file .env di root directory
	// dan memasukkan variabelnya ke environment
	if err := godotenv.Load(); err != nil {
		return nil, err // Jika error, return error
	}

	// 2. Konversi DB_PORT dari string ke integer
	// strconv.Atoi = string to integer
	// getEnv("DB_PORT", "5432") artinya:
	// - Cari environment variable DB_PORT
	// - Jika tidak ada, gunakan default "5432"
	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	// 3. Return struct Config dengan semua nilai
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     port, // sudah integer
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "auth_rbac"),
		JWTSecret:  getEnv("JWT_SECRET", "default-secret-key"),
		Port:       getEnv("PORT", "8080"),
	}, nil
}

// getEnv helper function untuk membaca environment variable
func getEnv(key, defaultValue string) string {
	// os.LookupEnv mencari environment variable
	// Return value, exists
	if value, exists := os.LookupEnv(key); exists {
		return value // Jika ada, return nilainya
	}
	return defaultValue // Jika tidak ada, return default
}
