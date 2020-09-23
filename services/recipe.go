package services

import (
	"context"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/util/expand"
	"github.com/Rahul12344/Recipes/util/parsing"
	"github.com/Rahul12344/Recipes/util/uuid"
	"github.com/Rahul12344/skelego"
)

//RecipeStore RecipeStore.
type RecipeStore interface {
	AddRecipe(recipe *models.Recipe)
	AddIngredients(ingredients ...*models.Ingredients)
	AddInstructions(instructions ...*models.Instructions)
	AddQuantities(quantities ...*models.Quantities)
	AddTags(tags ...*models.Tags)
	AddRecipeIngredients(recipeID string, ingredients []*models.Ingredients, quantities []*models.Quantities, length int)
	FindRecipe(matches []*models.Ingredients) []*models.Recipe
}

//RecipeIndex index for recipe
type RecipeIndex interface {
	GetRecipes(ctx context.Context, index, qType string, conditional map[string]interface{}, offset, limit int) []models.TotalRecipe
	QueryBuilder(ingredients ...string) map[string]interface{}
}

//RecipeService RecipeService.
type RecipeService struct {
	recipeStore RecipeStore
	index       RecipeIndex
	logger      skelego.Logging
}

//NewRecipeService constructs NewRecipeService.
func NewRecipeService(recipeStore RecipeStore, index RecipeIndex, logger skelego.Logging) *RecipeService {
	return &RecipeService{
		recipeStore: recipeStore,
		index:       index,
		logger:      logger,
	}
}

//Search Searches ES Index for corresponding recipe
func (rs *RecipeService) Search(ctx context.Context, ingredients []string) []models.TotalRecipe {
	return rs.index.GetRecipes(ctx, "recipes", "recipe", rs.index.QueryBuilder(ingredients...), 0, 20)
}

//Find finds matching recipes from ingredients.
func (rs *RecipeService) Find(ingredients []string) []models.TotalRecipe {

	/* TODO - load dataset */
	var dataset []string

	associations := expand.NewAssociations(ingredients, dataset)
	associations.Associate(0.2)

	for _, root := range associations.Roots {
		var ingredients []*models.Ingredients
		for _, ingredient := range root.AssociatedWords {
			ingredients = append(ingredients, &models.Ingredients{
				Ingredient: ingredient,
			})
		}
	}

	return nil
}

//NewRecipe creates new recipe.
func (rs *RecipeService) NewRecipe(recipeName, description string, numIngredients int, cc models.CountryCode, ingredientList []models.IngredientsTie) {
	recipeID := uuid.UUID()
	rs.recipeStore.AddRecipe(&models.Recipe{
		RecipeID:       recipeID,
		Name:           recipeName,
		Description:    description,
		NumIngredients: numIngredients,
		Country:        cc,
	})
	ingredients, quantities := models.MakeIngredientsAndQuantities(ingredientList...)
	rs.recipeStore.AddIngredients(ingredients...)
	rs.recipeStore.AddQuantities(quantities...)
	rs.recipeStore.AddRecipeIngredients(recipeID, ingredients, quantities, numIngredients)
}

//Image creates recipe model from an image.
func (rs *RecipeService) Image(filename string) []*models.Ingredients {
	parsing := parsing.NewParser()
	ingredients := parsing.Deconstruct(filename)

	var search []*models.Ingredients
	for _, ingredient := range ingredients {
		search = append(search, &models.Ingredients{
			Ingredient: ingredient,
		})
	}
	return search
}
