package models

type User struct {
	ID       string `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type Time struct {
	Value int    `json:"value" validate:"required,min:1,max:99999"`
	Unit  string `json:"unit" validate:"required,maxStringLength:10"`
}

type Ingredient struct {
	Name     string `json:"name" validate:"required,maxStringLength:50"`
	Quantity int    `json:"quantity" validate:"required,min:1,max:99999"`
	Unit     string `json:"unit" validate:"required,maxStringLength:10"`
}

type IngredientsGroup struct {
	Title       string       `json:"title" validate:"required,maxStringLength:50"`
	Ingredients []Ingredient `json:"ingredients" validate:"required,minArrayLength:1"`
}

type Method struct {
	Text string `json:"text" validate:"required,maxStringLength:500"`
}

type MethodsGroup struct {
	Title   string   `json:"title" validate:"required,maxStringLength:50"`
	Methods []Method `json:"methods" validate:"required,minArrayLength:1"`
}

type Recipe struct {
	ID                string             `json:"id"`
	Title             string             `json:"title" validate:"required,maxStringLength:50"`
	Image             string             `json:"image" validate:"required,maxStringLength:200"`
	Time              Time               `json:"time" validate:"required"`
	IngredientsGroups []IngredientsGroup `json:"ingredientsGroups" validate:"required,minArrayLength:1"`
	MethodsGroups     []MethodsGroup     `json:"methodsGroups" validate:"required,minArrayLength:1"`
	Author            User               `json:"author"`
}
