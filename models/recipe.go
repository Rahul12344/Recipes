package models

import (
	"github.com/Rahul12344/Recipes/util/uuid"
	"github.com/jinzhu/gorm"
)

//TotalRecipe all JSON
type TotalRecipe struct {
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	NumIngredients int              `json:"num_ingredients"`
	Country        CountryCode      `json:"country"`
	Ingredients    []IngredientsTie `json:"ingredients"`
	Instructions   []string         `json:"instructions"`
	Tags           []Tags           `json:"tags"`
}

// Recipe recipe model
type Recipe struct {
	gorm.Model
	RecipeID       string      `json:"recipe_id" gorm:"unique;not null;primary_key"`
	Name           string      `json:"name" gorm:"not null"`
	Description    string      `json:"description" gorm:"not null"`
	NumIngredients int         `json:"num_ingredients" gorm:"not null"`
	Country        CountryCode `json:"country" gorm:"not null"`
}

//TableName name of table
func (Recipe) TableName() string {
	return "recipes"
}

//RecipeIngredients All recipe ingredients.
type RecipeIngredients struct {
	gorm.Model
	RecipeID     string `json:"recipe_id" gorm:"not null"`
	IngredientID string `json:"ingredient_id" gorm:"primary_key;not null"`
	QuantityID   string `json:"quantity_id"  gorm:"primary_key;not null"`
}

//IngredientsTie Ties ingredients and quantities.
type IngredientsTie struct {
	Ingredient string `json:"ingredient"`
	Quantity   string `json:"quantity"`
}

//MakeIngredientsAndQuantities Creates ingredients and tags
func MakeIngredientsAndQuantities(ingredientsAndQuantities ...IngredientsTie) ([]*Ingredients, []*Quantities) {
	var ingredients []*Ingredients
	var quantities []*Quantities
	for _, ingredientsAndQuantity := range ingredientsAndQuantities {
		ingredients = append(ingredients, &Ingredients{
			IngredientID: uuid.UUID(),
			Ingredient:   ingredientsAndQuantity.Ingredient,
		})
		quantities = append(quantities, &Quantities{
			QuantityID: uuid.UUID(),
			Quantity:   ingredientsAndQuantity.Quantity,
		})
	}
	return ingredients, quantities
}

//TableName name of table
func (RecipeIngredients) TableName() string {
	return "recipe_ingredients"
}

//Tags Tags
type Tags struct {
	TagID    string `json:"tag_id;unique" gorm:"unique;not null;primary_key"`
	RecipeID string `json:"recipe_id" gorm:"not null"`
	Tag      string `json:"tag" gorm:"not null"`
}

//MakeTags creates tags model
func MakeTags(tags ...string) []*Tags {
	var search []*Tags
	for _, tag := range tags {
		search = append(search, &Tags{
			Tag: tag,
		})
	}
	return search
}

//TableName name of table
func (Tags) TableName() string {
	return "tags"
}

//Ingredients all recipe ingredients
type Ingredients struct {
	gorm.Model
	IngredientID string `json:"ingredient_id;unique" gorm:"unique;not null;primary_key"`
	Ingredient   string `json:"ingredient" gorm:"not null;unique"`
	Image        string `json:"image" gorm:"not null"`
}

//MakeIngredients Create ingredients model
func MakeIngredients(ingredients ...string) []*Ingredients {
	var search []*Ingredients
	for _, ingredient := range ingredients {
		search = append(search, &Ingredients{
			Ingredient: ingredient,
		})
	}
	return search
}

//TableName name of table
func (Ingredients) TableName() string {
	return "ingredients"
}

//Instructions all instructions
type Instructions struct {
	gorm.Model
	InstructionID  string `json:"instruction_id;unique" gorm:"unique;not null;primary_key"`
	RecipeID       string `json:"recipe_id" gorm:"not null"`
	Instruction    string `json:"instruction" gorm:"not null"`
	InstructionNum string `json:"instruction_num" gorm:"not null"`
}

//MakeInstructions create instruction model
func MakeInstructions(instructions []string, instructionNums []string) []*Instructions {
	var search []*Instructions
	for i := 0; i < len(instructions); i++ {
		search = append(search, &Instructions{
			Instruction:    instructions[i],
			InstructionNum: instructionNums[i],
		})
	}
	return search
}

//TableName name of table
func (Instructions) TableName() string {
	return "instructions"
}

//Quantities all quantities
type Quantities struct {
	gorm.Model
	QuantityID string `json:"quantity_id" gorm:"unique;not null;primary_key"`
	Quantity   string `json:"quantity" gorm:"not null"`
}

//MakeQuantities create quantities model
func MakeQuantities(quantities ...string) []*Quantities {
	var search []*Quantities
	for _, quantity := range quantities {
		search = append(search, &Quantities{
			Quantity: quantity,
		})
	}
	return search
}

//TableName name of table
func (Quantities) TableName() string {
	return "quantities"
}
