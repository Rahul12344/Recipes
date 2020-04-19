package postgres

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/jinzhu/gorm"
)

// UserRecipeStore handles recipe user interaction
type UserRecipeStore struct {
	client *gorm.DB
}

// NewUserRecipeStore Postgresql client
func NewUserRecipeStore(client *gorm.DB) *UserRecipeStore {
	const SchemaQuery = `CREATE SCHEMA IF NOT EXISTS users`
	client.Exec(SchemaQuery)
	//client.Exec(`set search_path='users'`)
	return &UserRecipeStore{
		client: client,
	}
}

func (urs *UserRecipeStore) create() {
	urs.client.AutoMigrate(&models.UserRecipes{})
}

//ADD adds recipe to user
func (urs *UserRecipeStore) ADD(userID string, recipe *models.Recipe) {
	urs.client.Create(models.UserRecipes{
		Adder:   userID,
		Recipes: recipe.RecipeID,
		Hits:    0,
	})
}

//REMOVE removes recipe from user
func (urs *UserRecipeStore) REMOVE(userID string, recipe *models.Recipe) {
	urs.client.Delete(models.UserRecipes{
		Adder: userID,
	})
}
