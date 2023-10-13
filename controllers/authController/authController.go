package authcontroller

import (
	"api-toko/helper"
	"api-toko/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	_, err := c.Cookie("refreshToken")
	if err == nil {
		c.JSON(400, gin.H{
			"message": "anda sudah login",
		})
		return
	}

	var user models.User

	type InputData struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var input InputData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
			"error":   err.Error(),
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("field %s %s", field, tag),
			})
			return
		}
	}

	if models.DB.Where("username=?", input.Username).First(&user).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "Username dan password tidak sesuai",
		})
		return
	}

	hashedPassword := []byte(user.Password)

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(input.Password)); err != nil {
		c.JSON(400, gin.H{
			"message": "Username dan password tidak sesuai",
		})
		return
	}

	token_refresh, _ := helper.GenerateTokenRefresh(int(user.ID))
	token_akses, _ := helper.GenerateTokenAkses(int(user.ID))

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    token_refresh,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	models.DB.Model(&user).Where("id=?", user.ID).Updates(map[string]interface{}{
		"refresh_token": token_refresh,
	})

	c.JSON(200, gin.H{
		"message":      "login berhasil",
		"access_token": token_akses,
	})
}

func Registrasi(c *gin.Context) {
	_, err := c.Cookie("refreshToken")
	if err == nil {
		c.JSON(400, gin.H{
			"message": "silahkan logout terlebih dahulu",
		})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
			"error":   err.Error(),
		})
		return
	}

	listError := []map[string]string{}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			if tag == "min" {
				tag = "min 5 character"
			}
			listError = append(listError, map[string]string{
				field: fmt.Sprintf("field %s %s", field, tag),
			})
		}
	}

	if len(listError) > 0 {
		c.JSON(400, gin.H{
			"message": "data tidak sesuai",
			"error":   listError,
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "gagal",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "registrasi berhasil",
	})
}

func Logout(c *gin.Context) {
	var user models.User
	refreshToken, _ := c.Cookie("refreshToken")

	models.DB.Model(&user).Where("refresh_token=?", refreshToken).Updates(map[string]interface{}{
		"refresh_token": nil,
	})

	expiredCookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, expiredCookie)

	c.JSON(200, gin.H{
		"message": "Logout berhasil",
	})
}

func RefreshToken(c *gin.Context) {
	refreshToken, _ := c.Cookie("refreshToken")

	var user models.User

	if models.DB.Where("refresh_token=?", refreshToken).First(&user).RowsAffected == 0 {
		c.JSON(403, gin.H{
			"message": "refresh token tidak ditemukan",
		})
		return
	}

	token_akses, _ := helper.GenerateTokenAkses(int(user.ID))
	c.JSON(200, gin.H{
		"access_token": token_akses,
	})

}

func GetFile(c *gin.Context) {
	filename := c.DefaultQuery("path", "")
	if filename == "" {
		c.JSON(404, gin.H{"message": "Parameter path kosong"})
		return
	}

	filename = filepath.Join("public", filename)

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		c.JSON(404, gin.H{"message": "Berkas tidak ditemukan"})
		return
	}

	c.File(filename)
}
