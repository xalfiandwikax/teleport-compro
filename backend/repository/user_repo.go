package repository

import (
	"backend/server/internal/models"

	"gorm.io/gorm"
)

// UserRepository interface mendefinisikan contract
// Interface ini berguna untuk testing dan abstraction
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
}

// userRepository implementasi konkrit dari interface
type userRepository struct {
	db *gorm.DB // Menyimpan koneksi database
}

// NewUserRepository factory function untuk membuat repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create menyimpan user baru ke database
func (r *userRepository) Create(user *models.User) error {
	// db.Create melakukan INSERT INTO users ...
	return r.db.Create(user).Error // Return error jika ada
}

// FindByEmail mencari user berdasarkan email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User // Variabel untuk menampung hasil

	// db.Where("email = ?", email).First(&user)
	// SQL: SELECT * FROM users WHERE email = 'email' LIMIT 1
	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err // Return pointer ke user dan error
}

// FindByID mencari user berdasarkan ID
func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User

	// db.First(&user, id)
	// SQL: SELECT * FROM users WHERE id = id LIMIT 1
	err := r.db.First(&user, id).Error

	return &user, err
}
