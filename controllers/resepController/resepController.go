package resepcontroller

import (
	"api-toko/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Index(c *gin.Context) {
	var resep []models.Resep

	models.DB.Preload("Kategori").Find(&resep)

	c.JSON(200, gin.H{
		"data": resep,
	})
}

func Show(c *gin.Context) {
	var resep models.Resep
	id := c.Param("id")

	if models.DB.Preload("Kategori").Find(&resep, id).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "data tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": resep,
	})
}

func Create(c *gin.Context) {
	var resep models.Resep

	if err := c.ShouldBindJSON(&resep); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
		})
		return
	}

	listError := []map[string]string{}
	validate := validator.New()
	if err := validate.Struct(resep); err != nil {
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

	models.DB.Create(&resep)
	c.JSON(200, gin.H{
		"data": resep,
	})
}

func Update(c *gin.Context) {
	var resep models.Resep
	id := c.Param("id")

	if err := c.ShouldBindJSON(&resep); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
		})
		return
	}

	listError := []map[string]string{}
	validate := validator.New()
	if err := validate.Struct(resep); err != nil {
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

	if models.DB.Model(&resep).Where("id=?", id).Updates(&resep).RowsAffected == 0 {
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
	var resep models.Resep
	id := c.Param("id")

	if models.DB.Delete(&resep, id).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "gagal menghapus data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "data berhasil dihapus",
	})
}
