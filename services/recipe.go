package services

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/util/expand"
	"github.com/Rahul12344/Recipes/util/parsing"
)

//RecipeStore recipe store
type RecipeStore interface {
	FIND(matches []*models.Ingredients) []*models.Recipe
	INGREDIENTS(ingredients []string) []*models.Ingredients
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
func (rs *RecipeService) FIND(ingredients []string) [][]*models.Recipe {

	/* TODO - load dataset */
	var dataset []string

	associations := expand.NewAssociations(ingredients, dataset)
	associations.Associate(0.2)

	var recipes [][]*models.Recipe
	for _, root := range associations.Roots {
		var ingredients []*models.Ingredients
		for _, ingredient := range root.AssociatedWords {
			ingredients = append(ingredients, &models.Ingredients{
				Ingredient: ingredient,
			})
		}
		recipes = append(recipes, rs.recipeStore.FIND(ingredients))
	}

	return recipes
}

//INGREDIENTS create recipe model
func (rs *RecipeService) INGREDIENTS(ingredients []string) []*models.Ingredients {
	return rs.recipeStore.INGREDIENTS(ingredients)
}

//IMAGE create recipe model
func (rs *RecipeService) IMAGE(filename string) []*models.Ingredients {
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

//CREATEDELIVERY find location of closest ingredients
func (rs *RecipeService) CREATEDELIVERY(ingredients ...string) {

}

//FINDLOCATION find location of closest ingredients
func (rs *RecipeService) FINDLOCATION(ingredients ...string) {
	for _, ingredient := range ingredients {
		rs.findClosestStoreAtCheapestPrice(ingredient)
	}
}

func (rs *RecipeService) findClosestStoreAtCheapestPrice(ingredient string) map[string]interface{} {
	return nil
}
