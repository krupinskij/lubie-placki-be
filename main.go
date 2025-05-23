package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lubie-placki-be/middlewares"
	"github.com/lubie-placki-be/routes"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.Headers())
	router.Use(middlewares.Github())

	routes.AuthRoutes(router)
	routes.ImageRoutes(router)
	routes.RecipeRoutes(router)

	router.Run("localhost:8080")
}
