package main

import (
	"github.com/lubie-placki-be/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RecipeRoutes(router)

	router.Run("localhost:8080")
}
