package postgres

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

// NewRecipeStore inits Recipe
func NewRecipeStore(client sqlservice.ORMWrapper) *RecipeStore {
	return &RecipeStore{
		client: client,
	}
}

//AddRecipe Creates new recipe.
func (rs *RecipeStore) AddRecipe(recipe *models.Recipe) {
	rs.client.ORM().Create(recipe)
}

//AddIngredients Adds ingredients to database
func (rs *RecipeStore) AddIngredients(ingredients ...*models.Ingredients) {
	for _, ingredient := range ingredients {
		if rs.client.ORM().NewRecord(ingredient) {
			rs.client.ORM().Create(ingredient)
		}
	}
}

//AddInstructions Adds instructions to database
func (rs *RecipeStore) AddInstructions(instructions ...*models.Instructions) {
	for _, instruction := range instructions {
		rs.client.ORM().Create(instruction)
	}
}

//AddQuantities Adds quantities to database
func (rs *RecipeStore) AddQuantities(quantities ...*models.Quantities) {
	for _, quantity := range quantities {
		rs.client.ORM().Create(quantity)
	}
}

//AddTags Adds tags to database
func (rs *RecipeStore) AddTags(tags ...*models.Tags) {
	for _, tag := range tags {
		rs.client.ORM().Create(tag)
	}
}

//AddRecipeIngredients Adds recipe ingredients to database
func (rs *RecipeStore) AddRecipeIngredients(recipeID string, ingredients []*models.Ingredients, quantities []*models.Quantities, length int) {
	for index := 0; index < length; index++ {
		recipeIngredient := &models.RecipeIngredients{
			RecipeID:     recipeID,
			IngredientID: ingredients[index].IngredientID,
			QuantityID:   quantities[index].QuantityID,
		}
		rs.client.ORM().Create(recipeIngredient)
	}
}

// FindRecipe Finds matching recipes based on ingredient from Postgres.
func (rs *RecipeStore) FindRecipe(matches []*models.Ingredients) []*models.Recipe {
	currentRelation := rs.client.ORM().Table("recipe_ingredients")
	for _, match := range matches {
		currentRelation.Where("IngredientID = ?", match.IngredientID)
	}
	return nil
}
