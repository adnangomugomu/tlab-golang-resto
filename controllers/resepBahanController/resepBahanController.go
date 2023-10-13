package resepbahancontroller

import (
	"api-toko/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Index(c *gin.Context) {
	var resep_bahan []models.ResepBahan

	models.DB.Preload("Resep.Kategori").Preload("Bahan").Find(&resep_bahan)

	c.JSON(200, gin.H{
		"data": resep_bahan,
	})
}

func Show(c *gin.Context) {
	var resep_bahan models.ResepBahan
	id := c.Param("id")

	if models.DB.Preload("Resep.Kategori").Preload("Bahan").Find(&resep_bahan, id).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "data tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": resep_bahan,
	})
}

func Create(c *gin.Context) {
	var resep_bahan models.ResepBahan

	if err := c.ShouldBindJSON(&resep_bahan); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
		})
		return
	}

	listError := []map[string]string{}
	validate := validator.New()
	if err := validate.Struct(resep_bahan); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
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

	models.DB.Create(&resep_bahan)
	c.JSON(200, gin.H{
		"data": resep_bahan,
	})
}

func Update(c *gin.Context) {
	var resep_bahan models.ResepBahan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&resep_bahan); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
		})
		return
	}

	listError := []map[string]string{}
	validate := validator.New()
	if err := validate.Struct(resep_bahan); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
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

	if models.DB.Model(&resep_bahan).Where("id=?", id).Updates(&resep_bahan).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "gagal update data",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "data berhasil diperbarui",
	})
}

func Delete(c *gin.Context) {
	var resep_bahan models.ResepBahan
	id := c.Param("id")

	if models.DB.Delete(&resep_bahan, id).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "gagal menghapus data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "data berhasil dihapus",
	})
}
