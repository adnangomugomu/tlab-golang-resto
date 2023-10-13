package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckCookies(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")

	if err != nil {
		c.JSON(403, gin.H{
			"message": "Akses tidak diizinkan",
		})
		c.Abort()
		return
	}

	jwtKey := os.Getenv("JWT_KEY_REFRESH")

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		c.JSON(401, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	if token.Valid {
		c.Next()
	} else {
		c.JSON(401, gin.H{"message": "Token is not valid"})
		c.Abort()
		return
	}
}
