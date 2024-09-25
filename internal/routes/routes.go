package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/lib", GetLib)
	rg := r.Group("/lib/song")
	rg.DELETE("/", DeleteSong)
	rg.PATCH("/", UpdateSong)
	rg.POST("/", CreateSong)
	rg.GET("/", GetSong)
	rg.GET("/text", GetText)

}