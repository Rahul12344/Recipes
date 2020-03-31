package services

import "github.com/Rahul12344/Recipes/models"

//RecipeStore recipe store
type RecipeStore interface {
	FIND(image bool, matches *models.Recipe) []*models.Recipe
	INGREDIENTS(ingredients []string) *models.Recipe
	IMAGE(filename string) *models.Recipe
}

//RecipeService recipe service
type RecipeService struct {
	recipeStore RecipeStore
}

//NewRecipeService constructs new recipe service
func NewRecipeService(recipeStore RecipeStore) *RecipeService {
	return &RecipeService{
		recipeStore: recipeStore,
	}
}

//FIND finds matching recipes
func (rs *RecipeService) FIND(image bool, matches *models.Recipe) []*models.Recipe {
	return rs.recipeStore.FIND(image, matches)
}

//INGREDIENTS create recipe model
func (rs *RecipeService) INGREDIENTS(ingredients []string) *models.Recipe {
	return rs.recipeStore.INGREDIENTS(ingredients)
}

//IMAGE create recipe model
func (rs *RecipeService) IMAGE(filename string) *models.Recipe {
	return rs.recipeStore.IMAGE(filename)
}
