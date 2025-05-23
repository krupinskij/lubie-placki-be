package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lubie-placki-be/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/auth/login", controllers.Login)
	router.GET("/auth/logout", controllers.Logout)
	router.GET("/auth/callback", controllers.Callback)
	router.GET("/auth/me", controllers.GetMe)
}
