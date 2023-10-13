package kategoricontroller

import (
	"api-toko/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Index(c *gin.Context) {
	var kategori []models.Kategori

	models.DB.Find(&kategori)

	c.JSON(200, gin.H{
		"data": kategori,
	})
}

func Show(c *gin.Context) {
	var kategori models.Kategori
	id := c.Param("id")

	if models.DB.Find(&kategori, id).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "data tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": kategori,
	})
}

func Create(c *gin.Context) {
	var kategori models.Kategori

	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
		})
		return
	}

	listError := []map[string]string{}
	validate := validator.New()
	if err := validate.Struct(kategori); err != nil {
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

	models.DB.Create(&kategori)
	c.JSON(200, gin.H{
		"data": kategori,
	})
}

func Update(c *gin.Context) {
	var kategori models.Kategori
	id := c.Param("id")

	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
		})
		return
	}

	listError := []map[string]string{}
	validate := validator.New()
	if err := validate.Struct(kategori); err != nil {
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

	if models.DB.Model(&kategori).Where("id=?", id).Updates(&kategori).RowsAffected == 0 {
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
	var kategori models.Kategori
	id := c.Param("id")

	if models.DB.Delete(&kategori, id).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "gagal menghapus data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "data berhasil dihapus",
	})
}
