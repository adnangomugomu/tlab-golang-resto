package helper

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenRefresh(userID int) (string, error) {
	jwtKey := os.Getenv("JWT_KEY_REFRESH")

	secretKey := []byte(jwtKey)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Waktu kadaluarsa dalam detik (contoh: 24 jam)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenAkses(userID int) (string, error) {
	jwtKey := os.Getenv("JWT_KEY_AKSES")

	secretKey := []byte(jwtKey)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 20).Unix(), // Waktu kadaluarsa dalam detik (contoh: 24 jam)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
