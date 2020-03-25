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
	return &RecipeStore{
		client: client,
	}
}

func (rs *RecipeStore) create() {
	/* TODO: Maybe change migration model to maybe define DB relationships */
	rs.client.AutoMigrate(&models.Recipe{})
}

// FIND finds matching recipes
func (rs *RecipeStore) FIND(matches *models.Recipe) []*models.Recipe {
	return nil
}
