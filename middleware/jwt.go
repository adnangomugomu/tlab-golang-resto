package middleware

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Jwt(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(403, gin.H{"message": "Token not provided"})
		c.Abort()
		return
	}

	tokenString = strings.Split(tokenString, "Bearer ")[1]
	jwtKey := os.Getenv("JWT_KEY_AKSES")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		c.JSON(401, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	if token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("user_id", claims["user_id"].(float64))
		c.Next()
	} else {
		c.JSON(401, gin.H{"message": "Token is not valid"})
		c.Abort()
		return
	}
}
