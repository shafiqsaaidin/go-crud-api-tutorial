package recipes

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Represent individual ingredients
type Ingredient struct {
	Name string `json:"name"`
}
