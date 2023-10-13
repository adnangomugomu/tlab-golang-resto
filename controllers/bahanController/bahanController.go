package bahancontroller

import (
	"api-toko/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Index(c *gin.Context) {
	var bahan []models.Bahan

	models.DB.Find(&bahan)

	c.JSON(200, gin.H{
		"data": bahan,
	})
}

func Show(c *gin.Context) {
	var bahan models.Bahan
	id := c.Param("id")

	if models.DB.Find(&bahan, id).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "data tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": bahan,
	})
}

func Create(c *gin.Context) {
	var bahan models.Bahan

	if err := c.ShouldBindJSON(&bahan); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
		})
		return
	}

	listError := []map[string]string{}
	validate := validator.New()
	if err := validate.Struct(bahan); err != nil {
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

	models.DB.Create(&bahan)
	c.JSON(200, gin.H{
		"data": bahan,
	})
}

func Update(c *gin.Context) {
	var bahan models.Bahan
	id := c.Param("id")

	if err := c.ShouldBindJSON(&bahan); err != nil {
		c.JSON(400, gin.H{
			"message": "data tidak valid",
		})
		return
	}

	listError := []map[string]string{}
	validate := validator.New()
	if err := validate.Struct(bahan); err != nil {
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

	if models.DB.Model(&bahan).Where("id=?", id).Updates(&bahan).RowsAffected == 0 {
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
	var bahan models.Bahan
	id := c.Param("id")

	if models.DB.Delete(&bahan, id).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "gagal menghapus data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "data berhasil dihapus",
	})
}
