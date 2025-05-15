package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type time struct {
	Value int16  `json:"value"`
	Unit  string `json:"unit"`
}

type recipe struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Time    time   `json:"time"`
	Content string `json:"content"`
	Author  user   `json:"author"`
}

var me = user{ID: "1", Username: "krupinskij"}
var recipes = []recipe{
	{
		ID:      "1",
		Title:   "Murzynek",
		Image:   "https://cdn.aniagotuje.com/pictures/articles/2018/03/104896-v-1000x1000.jpg",
		Time:    time{Value: 180, Unit: "m"},
		Content: "Upiecz murzynka",
		Author:  me,
	},
	{
		ID:      "2",
		Title:   "Piernik",
		Image:   "https://wszystkiegoslodkiego.pl/storage/images/202110/piernik-weganski.jpg",
		Time:    time{Value: 250, Unit: "m"},
		Content: "Upiecz piernik",
		Author:  me,
	},
	{
		ID:      "3",
		Title:   "Sernik",
		Image:   "https://cdn.aniagotuje.com/pictures/articles/2018/11/165653-v-1000x1000.jpg",
		Time:    time{Value: 210, Unit: "m"},
		Content: "Upiecz sernik",
		Author:  me,
	},
}

func getRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, recipes)
}

func main() {
	router := gin.Default()
	router.GET("/recipes", getRecipes)

	router.Run("localhost:8080")
}
