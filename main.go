package main

import (
	"fmt"

	"github.com/lubie-placki-be/validators"
)

func main() {
	// router := gin.Default()
	u := validators.User{Name: "Bob", Email: "bob@mycompany.com", Address: validators.Address{Street: "Waska"}, Animals: []validators.Animal{
		{Breed: "dog", Name: "Azor"},
		{Breed: "cat", Name: "Pussy"},
		{Breed: "turtle", Name: "Alexander"},
	}}

	if message, ok := validators.Validate(u); !ok {
		fmt.Println(message.Message)
	}

	// router.Use(middlewares.CORS())

	// routes.RecipeRoutes(router)

	// router.Run("localhost:8080")
}
