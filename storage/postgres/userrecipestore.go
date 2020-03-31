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
	return &UserRecipeStore{
		client: client,
	}
}

func (urs *UserRecipeStore) create() {
	/* TODO: Maybe change migration model to maybe define DB relationships */
	urs.client.AutoMigrate(&models.RecipePointer{})
}

//ADD adds recipe to user
func (urs *UserRecipeStore) ADD(uuid string, recipe *models.Recipe) {
	urs.client.Create(models.RecipePointer{
		Adder:  uuid,
		Recipe: recipe.UUID,
		Hits:   0,
	})
}

//REMOVE removes recipe from user
func (urs *UserRecipeStore) REMOVE(uuid string, recipe *models.Recipe) {
	urs.client.Delete(models.RecipePointer{
		Adder:  uuid,
		Recipe: recipe.UUID,
	})
}
