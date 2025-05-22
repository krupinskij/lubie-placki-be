package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lubie-placki-be/controllers"
)

func ImageRoutes(router *gin.Engine) {
	router.GET("/images/:id", controllers.DownloadImage)
	router.POST("/images", controllers.UploadImage)
}
