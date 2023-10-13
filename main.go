package main

import (
	"api-toko/middleware"
	"api-toko/models"
	"api-toko/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:9000"}
	r.Use(cors.New(config))

	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.Jwt)

	globalGroup := r.Group("/api")
	globalGroup.Use(middleware.CheckCookies)

	routers.AuthRouter(r)
	routers.GlobalRouter(globalGroup)
	routers.ProfilRouter(apiGroup)
	routers.RefKategoriRouter(apiGroup)
	routers.BahanRouter(apiGroup)
	routers.ResepRouter(apiGroup)
	routers.ResepBahanRouter(apiGroup)

	r.Run(":9000")
}
