package handlers

import (
	"backend/server/internal/models"
	repository "backend/server/internal/repo"
	"backend/server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler struct menyimpan dependencies
type AuthHandler struct {
	userRepo  repository.UserRepository // Untuk akses database
	jwtSecret string                    // Secret key untuk JWT
}

// NewAuthHandler factory function
func NewAuthHandler(userRepo repository.UserRepository, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// LoginRequest struct untuk request body
// Gin binding tags untuk validasi otomatis
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`    // Required dan format email
	Password string `json:"password" binding:"required,min=6"` // Required, min 6 karakter
}

// LoginResponse struct untuk response
type LoginResponse struct {
	Token string      `json:"token"` // JWT Token
	User  models.User `json:"user"`  // User data
}

// Login handler untuk endpoint POST /auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	// 1. Parse dan validasi request body
	var req LoginRequest

	// ShouldBindJSON:
	// - Parse JSON dari request body
	// - Validasi berdasarkan tags di struct
	// - Return error jika validasi gagal
	if err := c.ShouldBindJSON(&req); err != nil {
		// Gin otomatis mengembalikan error validasi
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return // Stop execution jika request invalid
	}

	// 2. Cari user berdasarkan email
	user, err := h.userRepo.FindByEmail(req.Email)
	if err != nil {
		// User tidak ditemukan atau error database
		// Untuk keamanan, kasus "user tidak ditemukan" dan "password salah"
		// dikembalikan dengan pesan yang sama
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// 3. Verifikasi password
	// user.CheckPassword membandingkan password input dengan hash di database
	if err := user.CheckPassword(req.Password); err != nil {
		// Password tidak cocok
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// 4. Generate JWT token
	token, err := utils.GenerateJWT(
		user.ID,           // user_id untuk claims
		user.Email,        // email untuk claims
		string(user.Role), // role untuk claims
		h.jwtSecret,       // secret key untuk signing
	)

	if err != nil {
		// Error generate token
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate authentication token",
		})
		return
	}

	// 5. Siapkan response
	// Hapus password dari response untuk keamanan
	user.Password = ""

	// 6. Kirim response sukses
	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
		User:  *user,
	})
}
