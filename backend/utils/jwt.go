package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims adalah custom claims untuk JWT
// Claims = data yang disimpan di dalam token
type Claims struct {
	UserID               uint   `json:"user_id"` // ID user
	Email                string `json:"email"`   // Email user
	Role                 string `json:"role"`    // Role user
	jwt.RegisteredClaims        // Embedded standard claims
}

// GenerateJWT membuat token JWT baru
func GenerateJWT(userID uint, email, role, secret string) (string, error) {
	// Buat claims (payload)
	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expired 24 jam
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // Waktu dibuat
			Issuer:    "backend-api",                                      // Penerbit token
		},
	}

	// Buat token dengan algoritma HS256
	// jwt.NewWithClaims(method, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token dengan secret key
	// token.SignedString([]byte(secret))
	return token.SignedString([]byte(secret))
}

// ValidateJWT memvalidasi token dan mengembalikan claims
func ValidateJWT(tokenString, secret string) (*Claims, error) {
	// Parse token dengan claims kita
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{}, // Target struct untuk parse
		func(token *jwt.Token) (interface{}, error) {
			// Validasi signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(secret), nil // Return secret key untuk validasi
		},
	)

	if err != nil {
		return nil, err // Token invalid/expired
	}

	// Extract claims dari token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil // Token valid, return claims
	}

	return nil, errors.New("invalid token")
}
