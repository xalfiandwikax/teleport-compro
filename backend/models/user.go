package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Role adalah tipe custom untuk role user
type Role string

// Konstanta untuk role
const (
	RoleAdmin      Role = "admin"
	RoleSuperadmin Role = "superadmin"
)

// User struct merepresentasikan tabel users di database
type User struct {
	gorm.Model        // Embedded struct dari GORM
	Name       string `gorm:"not null" json:"name"`         // Kolom tidak boleh null
	Email      string `gorm:"unique;not null" json:"email"` // Email harus unique
	Password   string `gorm:"not null" json:"-"`            // "-" artinya tidak di-export ke JSON
	Role       Role   `gorm:"type:varchar(20);not null" json:"role"`
}

// HashPassword meng-hash password sebelum disimpan ke database
func (u *User) HashPassword() error {
	// bcrypt.GenerateFromPassword menghasilkan hash dari password
	// bcrypt.DefaultCost = 10 (tingkat kesulitan hashing)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword) // Simpan hash, bukan password plain text
	return nil
}

// CheckPassword membandingkan password input dengan hash di database
func (u *User) CheckPassword(password string) error {
	// bcrypt.CompareHashAndPassword membandingkan hash dengan password
	// Jika cocok return nil, jika tidak return error
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// BeforeCreate adalah hook GORM yang dijalankan sebelum create
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Pastikan role valid
	if u.Role != RoleAdmin && u.Role != RoleSuperadmin {
		return errors.New("invalid role")
	}
	return u.HashPassword() // Auto-hash password sebelum save
}
