package postgres

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/util/uuid"
	"github.com/jinzhu/gorm"
)

// RecipeStore contains db handling for recipe objects
type RecipeStore struct {
	client *gorm.DB
}

// NewRecipeStore inits Recipe
func NewRecipeStore(client *gorm.DB) *RecipeStore {
	return &RecipeStore{
		client: client,
	}
}

func (rs *RecipeStore) create() {
	/* TODO: Maybe change migration model to maybe define DB relationships */
	rs.client.AutoMigrate(&models.Recipe{}, &models.RecipeIngredients{},
		&models.Ingredients{}, &models.Instructions{}, &models.Quantities{},
		&models.Tags{})
}

//INSERT creates new recipe
func (rs *RecipeStore) INSERT(recipe *models.Recipe, ingredients []*models.Ingredients, instructions []*models.Instructions, tags []*models.Tags, quantities []*models.Quantities) {
	rs.client.Create(recipe)
	for i := 0; i < len(ingredients); i++ {
		if err := rs.client.Where("Ingredient = ?", ingredients[i].Ingredient).Find(&models.Ingredients{}); err != nil {
			ingredients[i].IngredientID = uuid.UUID()
			rs.client.Create(ingredients[i])
		}
		if err := rs.client.Where("Quantity = ?", quantities[i].Quantity).Find(&models.Quantities{}); err != nil {
			quantities[i].QuantityID = uuid.UUID()
			rs.client.Create(quantities[i])
		}

		rs.client.Model(&recipe).Association("RecipeIngredients").Append(&models.RecipeIngredients{
			IngredientID: ingredients[i].IngredientID,
			QuantityID:   quantities[i].QuantityID,
		})

	}
}

// FIND finds matching recipes
func (rs *RecipeStore) FIND(matches []*models.Ingredients) []*models.Recipe {
	//var recipes []*models.RecipeIngredients

	currentRelation := rs.client.Table("recipe_ingredients")
	for _, match := range matches {
		currentRelation.Where("IngredientID = ?", match.IngredientID)
	}
	return nil
}

//INGREDIENTS create recipe model
func (rs *RecipeStore) INGREDIENTS(ingredients []string) []*models.Ingredients {
	var search []*models.Ingredients
	for _, ingredient := range ingredients {
		search = append(search, &models.Ingredients{
			Ingredient: ingredient,
		})
	}
	return search
}
