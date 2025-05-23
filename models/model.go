package models

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	GithubID int64         `json:"githubId" bson:"github_id"`
	Login    string        `json:"login" bson:"login" validate:"required"`
	Name     string        `json:"name" bson:"name" validate:"required"`
}

type Time struct {
	Value int    `json:"value" bson:"value" validate:"required,min:1,max:99999"`
	Unit  string `json:"unit" bson:"unit" validate:"required,maxStringLength:10"`
}

type Ingredient struct {
	Name     string `json:"name" bson:"name" validate:"required,maxStringLength:50"`
	Quantity int    `json:"quantity" bson:"quantity" validate:"required,min:1,max:99999"`
	Unit     string `json:"unit" bson:"unit" validate:"required,maxStringLength:10"`
}

type IngredientsGroup struct {
	Title       string       `json:"title" bson:"title" validate:"required,maxStringLength:50"`
	Ingredients []Ingredient `json:"ingredients" bson:"ingredients" validate:"required,minArrayLength:1,deep"`
}

type Method struct {
	Text string `json:"text" bson:"text" validate:"required,maxStringLength:500"`
}

type MethodsGroup struct {
	Title   string   `json:"title" bson:"title" validate:"required,maxStringLength:50"`
	Methods []Method `json:"methods" bson:"methods" validate:"required,minArrayLength:1,deep"`
}

type Recipe struct {
	ID                bson.ObjectID      `json:"id" bson:"_id,omitempty"`
	Title             string             `json:"title" bson:"title" validate:"required,maxStringLength:50"`
	ImageId           string             `json:"imageId" bson:"image_id" validate:"required"`
	Time              Time               `json:"time" bson:"time" validate:"required,deep"`
	IngredientsGroups []IngredientsGroup `json:"ingredientsGroups" bson:"ingredients_groups" validate:"required,minArrayLength:1,deep"`
	MethodsGroups     []MethodsGroup     `json:"methodsGroups" bson:"methods_groups" validate:"required,minArrayLength:1,deep"`
	Author            User               `json:"author" bson:"author"`
}
