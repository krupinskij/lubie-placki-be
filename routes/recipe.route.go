package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lubie-placki-be/controllers"
)

func RecipeRoutes(router *gin.Engine) {
	router.GET("/recipes", controllers.GetAllRecipes)
	router.GET("/recipes/:id", controllers.GetRecipe)
	router.GET("/recipes/random", controllers.GetRandomId)
}
