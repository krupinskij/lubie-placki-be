package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubie-placki-be/configs"
	"github.com/lubie-placki-be/models"
	"github.com/lubie-placki-be/services"
)

func GetAllRecipes(c *gin.Context) {
	recipes, err := services.GetAllRecipes()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, recipes)
}

func GetRecipe(c *gin.Context) {
	id := c.Param("id")
	recipes, err := services.GetRecipeById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.Writer.Header().Set("Cache-Control", "max-age=3600, public")

	c.IndentedJSON(http.StatusOK, recipes)
}

func GetRandomId(c *gin.Context) {
	id, err := services.GetRandomId()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

func CreateRecipe(c *gin.Context) {
	var newRecipe models.Recipe

	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if message, ok := configs.Validate(newRecipe); !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": message.Message})
		return
	}

	id, err := services.CreateRecipe(newRecipe)

	if err != nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}
