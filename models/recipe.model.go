package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Time struct {
	Value int16  `json:"value"`
	Unit  string `json:"unit"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity int16  `json:"quantity"`
	Unit     string `json:"unit"`
}

type IngredientsGroup struct {
	Title       string       `json:"title"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Method struct {
	Text string `json:"text"`
}

type MethodsGroup struct {
	Title   string   `json:"title"`
	Methods []Method `json:"methods"`
}

type Recipe struct {
	ID                string             `json:"id"`
	Title             string             `json:"title" valid:"required,minstringlength(10),maxstringlength(50)"`
	Image             string             `json:"image"`
	Time              Time               `json:"time"`
	IngredientsGroups []IngredientsGroup `json:"ingredientsGroups"`
	MethodsGroups     []MethodsGroup     `json:"methodsGroups"`
	Author            User               `json:"author"`
}
