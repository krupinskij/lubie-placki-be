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

type Ingredients struct {
	Title       string       `json:"title"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Methods struct {
	Title   string   `json:"title"`
	Methods []string `json:"methods"`
}

type Recipe struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Image       string        `json:"image"`
	Time        Time          `json:"time"`
	Ingredients []Ingredients `json:"ingredients"`
	Methods     []Methods     `json:"methods"`
	Author      User          `json:"author"`
}
