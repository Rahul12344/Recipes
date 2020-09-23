package elasticsearch

import (
	"context"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/storage/elasticsearch/queries"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/index"
)

//NewRecipeIndex Create new Recipe index
func NewRecipeIndex(index index.Index, logger skelego.Logging) *RecipeIndex {
	return &RecipeIndex{
		index:  index,
		logger: logger,
	}
}

//AddRecipes Adds recipes to index
func (ri *RecipeIndex) AddRecipes(recipes ...models.TotalRecipe) error {
	return nil
}

//GetRecipes Searches on Recipe index
func (ri *RecipeIndex) GetRecipes(ctx context.Context, index, qType string, conditional map[string]interface{}, offset, limit int) []models.TotalRecipe {
	if !ri.verifySchema(conditional) {
		return nil
	}
	hits := ri.index.SearchIndex(ctx, "recipes", ri.index.Query(queries.CreateRecipeQueryForArgs(conditional), ri.logger), ri.logger)
	if hits == nil {
		return nil
	}
	return nil
}

//QueryBuilder Builds query
func (ri *RecipeIndex) QueryBuilder(ingredients ...string) map[string]interface{} {
	return nil
}

//VerifySchema Verifies query
func (ri *RecipeIndex) verifySchema(conditional map[string]interface{}) bool {
	return true
}
