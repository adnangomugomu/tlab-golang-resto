package profilcontroller

import (
	"api-toko/models"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GantiFoto(c *gin.Context) {
	user_id := c.MustGet("user_id").(float64)
	id := strconv.FormatFloat(user_id, 'f', -1, 64)

	file, err := c.FormFile("gambar")

	if err != nil {
		c.JSON(400, gin.H{
			"message": "gagal upload file",
			"error":   err.Error(),
		})
		return
	}

	var user models.User
	if models.DB.Where("id=?", id).Find(&user).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "Akun tidak ditemukan",
		})
		return
	}

	_ = os.Remove("public/" + user.FotoProfil)

	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := []string{".jpg", ".jpeg", ".png", ".webp"}

	valid := false
	for _, val := range allowedExts {
		if val == fileExt {
			valid = true
		}
	}

	if !valid {
		c.JSON(400, gin.H{
			"message": "jenis file tidak diizinkan",
		})
		return
	}

	random := id + "-" + strconv.Itoa(rand.Intn(10000)+1) + "-"
	fileSave := "foto-profil/" + random + file.Filename
	dst := "public/" + fileSave
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(500, gin.H{
			"message": "terjadi kesalahan",
			"error":   err.Error(),
		})
		return
	}

	if models.DB.Model(&user).Where("id=?", id).Updates(map[string]interface{}{
		"foto_profil": fileSave,
	}).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "gagal memperbarui data",
		})
		return
	}

	base_url := os.Getenv("BASE_URL")

	c.JSON(200, gin.H{
		"message": "data berhasil diperbarui",
		"file":    base_url + "file?path=" + fileSave,
	})
}
