package routers

import (
	authcontroller "api-toko/controllers/authController"
	bahancontroller "api-toko/controllers/bahanController"
	kategoricontroller "api-toko/controllers/kategoriController"
	profilcontroller "api-toko/controllers/profilController"
	resepbahancontroller "api-toko/controllers/resepBahanController"
	resepcontroller "api-toko/controllers/resepController"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	r.POST("api/auth/login", authcontroller.Login)
	r.POST("api/auth/registrasi", authcontroller.Registrasi)
	r.GET("file", authcontroller.GetFile)
}

func GlobalRouter(r *gin.RouterGroup) {
	r.GET("auth/refresh_token", authcontroller.RefreshToken)
	r.GET("auth/logout", authcontroller.Logout)
}

func ProfilRouter(r *gin.RouterGroup) {
	r.PATCH("profil/foto", profilcontroller.GantiFoto)
}

func RefKategoriRouter(r *gin.RouterGroup) {
	r.GET("kategori", kategoricontroller.Index)
	r.GET("kategori/:id", kategoricontroller.Show)
	r.POST("kategori", kategoricontroller.Create)
	r.PUT("kategori/:id", kategoricontroller.Update)
	r.DELETE("kategori/:id", kategoricontroller.Delete)
}

func BahanRouter(r *gin.RouterGroup) {
	r.GET("bahan", bahancontroller.Index)
	r.GET("bahan/:id", bahancontroller.Show)
	r.POST("bahan", bahancontroller.Create)
	r.PUT("bahan/:id", bahancontroller.Update)
	r.DELETE("bahan/:id", bahancontroller.Delete)
}

func ResepRouter(r *gin.RouterGroup) {
	r.GET("resep", resepcontroller.Index)
	r.GET("resep/:id", resepcontroller.Show)
	r.POST("resep", resepcontroller.Create)
	r.PUT("resep/:id", resepcontroller.Update)
	r.DELETE("resep/:id", resepcontroller.Delete)
}

func ResepBahanRouter(r *gin.RouterGroup) {
	r.GET("resep_bahan", resepbahancontroller.Index)
	r.GET("resep_bahan/:id", resepbahancontroller.Show)
	r.POST("resep_bahan", resepbahancontroller.Create)
	r.PUT("resep_bahan/:id", resepbahancontroller.Update)
	r.DELETE("resep_bahan/:id", resepbahancontroller.Delete)
}
