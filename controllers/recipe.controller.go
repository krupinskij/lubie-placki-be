package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lubie-placki-be/configs"
	"github.com/lubie-placki-be/middlewares"
	"github.com/lubie-placki-be/models"
	"github.com/lubie-placki-be/services"
)

func GetAllRecipes(c *gin.Context) {
	p := c.Query("page")
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 1
	}
	recipes, err := services.GetAllRecipes(page)

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
	if !middlewares.IsAuthenticated {
		c.IndentedJSON(http.StatusUnauthorized, nil)
	}

	var newRecipe models.Recipe

	if err := c.ShouldBindJSON(&newRecipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if result, ok := configs.Validate(newRecipe); !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": result.Message})
		return
	}

	id, err := services.CreateRecipe(newRecipe)

	if err != nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}
