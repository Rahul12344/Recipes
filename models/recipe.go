package models

import "github.com/jinzhu/gorm"

// Recipe recipe model
type Recipe struct {
	gorm.Model
	RecipeID       string              `json:"recipe_id" gorm:"unique;not null;primary_key"`
	Name           string              `json:"name"`
	Description    string              `json:"description"`
	NumIngredients int                 `json:"num_ingredients"`
	Ingredients    []RecipeIngredients `json:"recipe_ingredients" gorm:"many2many:ingredient_recipes"`
	Instructions   []Instructions      `json:"recipe_instructions" gorm:"many2many:recipe_instructions"`
	Tags           []Tags              `json:"recipe_tags" gorm:"many2many:recipe_tags"`
	Country        CountryCode         `json:"country"`
}

//TableName name of table
func (Recipe) TableName() string {
	return "recipe"
}

//RecipeIngredients all recipe ingredients
type RecipeIngredients struct {
	gorm.Model
	Recipes      []Recipe `json:"recipes" gorm:"many2many:ingredient_recipes"`
	IngredientID string   `json:"ingredient_id" gorm:"primary_key"`
	QuantityID   string   `json:"quantity_id"  gorm:"primary_key"`
}

//TableName name of table
func (RecipeIngredients) TableName() string {
	return "recipe_ingredients"
}

//Tags tags
type Tags struct {
	TagID   string   `json:"tag_id;unique" gorm:"unique;not null;primary_key"`
	Recipes []Recipe `json:"recipes" gorm:"many2many:recipe_tags"`
	Tag     string   `json:"tag"`
}

//TableName name of table
func (Tags) TableName() string {
	return "tags"
}

//Ingredients all recipe ingredients
type Ingredients struct {
	gorm.Model
	IngredientID string `json:"ingredient_id;unique" gorm:"unique;not null;primary_key"`
	Ingredient   string `json:"ingredient"`
	Image        string `json:"image"`
}

//TableName name of table
func (Ingredients) TableName() string {
	return "ingredients"
}

//Instructions all instructions
type Instructions struct {
	gorm.Model
	InstructionID  string   `json:"instruction_id;unique" gorm:"unique;not null;primary_key"`
	Recipes        []Recipe `json:"recipes" gorm:"many2many:recipe_instructions"`
	Instruction    string   `json:"instruction"`
	InstructionNum int      `json:"instruction_num"`
}

//TableName name of table
func (Instructions) TableName() string {
	return "instructions"
}

//Quantities all quantities
type Quantities struct {
	gorm.Model
	QuantityID string `json:"quantity_id" gorm:"unique;not null;primary_key"`
	Quantity   string `json:"quantity"`
}

//TableName name of table
func (Quantities) TableName() string {
	return "quantities"
}
