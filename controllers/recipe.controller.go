package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubie-placki-be/services"
)

func GetAllRecipes(c *gin.Context) {
	recipes, err := services.GetAllRecipes()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, recipes)
}

func GetRecipe(c *gin.Context) {
	id := c.Param("id")
	recipes, err := services.GetRecipeById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, recipes)
}
