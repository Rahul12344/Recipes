package postgres

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/jinzhu/gorm"
)

// RecipeStore contains db handling for recipe objects
type RecipeStore struct {
	client *gorm.DB
}

// NewRecipeStore inits Recipe
func NewRecipeStore(client *gorm.DB) *RecipeStore {
	const SchemaQuery = `CREATE SCHEMA IF NOT EXISTS recipes`
	client.Exec(SchemaQuery)
	client.Exec(`set search_path='recipes'`)
	return &RecipeStore{
		client: client,
	}
}

func (rs *RecipeStore) create() {
	/* TODO: Maybe change migration model to maybe define DB relationships */
	/*gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		return "recipes." + tableName
	}*/
	rs.client.AutoMigrate(&models.Recipe{}, &models.RecipeIngredients{}, &models.Ingredients{}, &models.Instructions{}, &models.Quantities{})
}

// FIND finds matching recipes
func (rs *RecipeStore) FIND(matches []*models.Ingredients) []*models.Recipe {
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
