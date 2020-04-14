package models

import "github.com/jinzhu/gorm"

// Recipe recipe model
type Recipe struct {
	gorm.Model
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Ingredients int         `json:"ingredients"`
	Tag         string      `json:"tag"`
	UUID        string      `json:"uuid"`
	Country     CountryCode `json:"country"`
}

//RecipeIngredients all recipe ingredients
type RecipeIngredients struct {
	gorm.Model
	UUID          string `json:"uuid"`
	IngredientID  string `json:"IngredientID"`
	InstructionID string `json:"InstructionID"`
	QuantityID    string `json:"QuantityID"`
}

//Ingredients all recipe ingredients
type Ingredients struct {
	gorm.Model
	IngredientID string `json:"IngredientID"`
	Ingredient   string `json:"Ingredient"`
	Image        string `json:"Image"`
}

//Instructions all instructions
type Instructions struct {
	gorm.Model
	InstructionID  string `json:"InstructionID"`
	Instruction    string `json:"Instruction"`
	InstructionNum int    `json:"InstructionNum"`
}

//Quantities all quantities
type Quantities struct {
	gorm.Model
	QuantityID string `json:"QuantityID"`
	Quantity   string `json:"Quantity"`
}
